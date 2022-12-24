package lbs

import (
	"fmt"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/twpayne/go-geom"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

type testAddress struct {
	ID      int
	Address *AddressEntity `gorm:"embedded"`
}

func TestAddress(t *testing.T) {
	var (
		dbName   = "mysql"
		user     = "root"
		password = "testPassword"
	)
	pool, err := dockertest.NewPool("")
	pool.MaxWait = time.Minute * 2
	require.NoError(t, err)

	resource, err := pool.Run("mysql", "5.7", []string{
		"MYSQL_ROOT_PASSWORD=" + password,
		"MYSQL_ROOT_HOST=%",
	})
	require.NoError(t, err)
	defer func() {
		assert.NoError(t, pool.Purge(resource))
	}()
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:%s)/%s?parseTime=true&loc=Local&charset=utf8mb4", user, password, resource.GetPort("3306/tcp"), dbName)
	var db *gorm.DB
	require.NoError(t, pool.Retry(func() error {
		var err error
		db, err = gorm.Open(mysql.Open(dsn))
		if err != nil {
			return err
		}
		return nil
	}))
	db = db.Debug()
	err = db.AutoMigrate(&testAddress{})
	assert.NoError(t, err)

	addr := &AddressEntity{
		Country: "Country",
		State:   "State",
		City:    "City",
		ZipCode: "ZipCode",
		Line1:   "Line1",
		Line2:   "Line2",
		Line3:   "Line3",
		Geo:     NewPoint(geom.NewPoint(geom.XY).MustSetCoords([]float64{0.1275, 51.50722}).SetSRID(4326)),
	}
	pbAddr, err := addr.ToPb()
	assert.NoError(t, err)
	addr, err = NewAddressEntityFromPb(pbAddr)
	assert.NoError(t, err)

	err = db.Create(&testAddress{ID: 1, Address: addr}).Error
	assert.NoError(t, err)

	var dbAddr testAddress
	err = db.Model(&testAddress{}).First(&dbAddr, "id = ?", 1).Error
	assert.NoError(t, err)
	assert.Equal(t, 1, dbAddr.ID)
	assert.NotNil(t, dbAddr.Address)
	assert.NotNil(t, dbAddr.Address.Geo)
	assert.Equal(t, "Country", dbAddr.Address.Country)
	assert.Equal(t, dbAddr.Address.Geo.Point.Point.SRID(), 4326)
	assert.Equal(t, dbAddr.Address.Geo.Point.Point.Coords().X(), 0.1275)
	assert.Equal(t, dbAddr.Address.Geo.Point.Point.Coords().Y(), 51.50722)

	err = db.Create(&testAddress{ID: 2, Address: &AddressEntity{
		Country: "Country",
		State:   "State",
		City:    "City",
		ZipCode: "ZipCode",
		Line1:   "Line1",
		Line2:   "Line2",
		Line3:   "Line3",
		Geo:     nil, // empty geo
	}}).Error
	assert.NoError(t, err)
	var dbAddr2 testAddress
	err = db.Model(&testAddress{}).First(&dbAddr2, "id = ?", 2).Error
	assert.NoError(t, err)
	assert.Equal(t, 2, dbAddr2.ID)
	assert.Nil(t, dbAddr2.Address.Geo)
}
