// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: address.proto

package lbs

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Country string `protobuf:"bytes,1,opt,name=country,proto3" json:"country,omitempty"`
	Region  string `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	// State or province
	State   string `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	City    string `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	ZipCode string `protobuf:"bytes,5,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
	Line1   string `protobuf:"bytes,6,opt,name=line1,proto3" json:"line1,omitempty"`
	Line2   string `protobuf:"bytes,7,opt,name=line2,proto3" json:"line2,omitempty"`
	Line3   string `protobuf:"bytes,8,opt,name=line3,proto3" json:"line3,omitempty"`
	// Geo geojson
	Geo *structpb.Struct `protobuf:"bytes,100,opt,name=geo,proto3" json:"geo,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_address_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_address_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_address_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Address) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Address) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetZipCode() string {
	if x != nil {
		return x.ZipCode
	}
	return ""
}

func (x *Address) GetLine1() string {
	if x != nil {
		return x.Line1
	}
	return ""
}

func (x *Address) GetLine2() string {
	if x != nil {
		return x.Line2
	}
	return ""
}

func (x *Address) GetLine3() string {
	if x != nil {
		return x.Line3
	}
	return ""
}

func (x *Address) GetGeo() *structpb.Struct {
	if x != nil {
		return x.Geo
	}
	return nil
}

var File_address_proto protoreflect.FileDescriptor

var file_address_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x6c, 0x62, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xed, 0x01, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x7a, 0x69,
	0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x7a, 0x69,
	0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x31, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6e, 0x65, 0x32, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x65,
	0x32, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x33, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x33, 0x12, 0x29, 0x0a, 0x03, 0x67, 0x65, 0x6f, 0x18, 0x64,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x03, 0x67,
	0x65, 0x6f, 0x42, 0x1c, 0x5a, 0x1a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x61, 0x61, 0x73, 0x2f, 0x6c, 0x62, 0x73, 0x3b, 0x6c, 0x62, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_address_proto_rawDescOnce sync.Once
	file_address_proto_rawDescData = file_address_proto_rawDesc
)

func file_address_proto_rawDescGZIP() []byte {
	file_address_proto_rawDescOnce.Do(func() {
		file_address_proto_rawDescData = protoimpl.X.CompressGZIP(file_address_proto_rawDescData)
	})
	return file_address_proto_rawDescData
}

var file_address_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_address_proto_goTypes = []interface{}{
	(*Address)(nil),         // 0: lbs.Address
	(*structpb.Struct)(nil), // 1: google.protobuf.Struct
}
var file_address_proto_depIdxs = []int32{
	1, // 0: lbs.Address.geo:type_name -> google.protobuf.Struct
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_address_proto_init() }
func file_address_proto_init() {
	if File_address_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_address_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_address_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_address_proto_goTypes,
		DependencyIndexes: file_address_proto_depIdxs,
		MessageInfos:      file_address_proto_msgTypes,
	}.Build()
	File_address_proto = out.File
	file_address_proto_rawDesc = nil
	file_address_proto_goTypes = nil
	file_address_proto_depIdxs = nil
}
