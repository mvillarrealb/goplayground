// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: geocoder.proto

package geocoder

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LatLng struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float32 `protobuf:"fixed32,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float32 `protobuf:"fixed32,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *LatLng) Reset() {
	*x = LatLng{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geocoder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LatLng) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LatLng) ProtoMessage() {}

func (x *LatLng) ProtoReflect() protoreflect.Message {
	mi := &file_geocoder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LatLng.ProtoReflect.Descriptor instead.
func (*LatLng) Descriptor() ([]byte, []int) {
	return file_geocoder_proto_rawDescGZIP(), []int{0}
}

func (x *LatLng) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *LatLng) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type GeoCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location *LatLng `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *GeoCodeRequest) Reset() {
	*x = GeoCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geocoder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeoCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoCodeRequest) ProtoMessage() {}

func (x *GeoCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_geocoder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoCodeRequest.ProtoReflect.Descriptor instead.
func (*GeoCodeRequest) Descriptor() ([]byte, []int) {
	return file_geocoder_proto_rawDescGZIP(), []int{1}
}

func (x *GeoCodeRequest) GetLocation() *LatLng {
	if x != nil {
		return x.Location
	}
	return nil
}

type GeocodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address  string  `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Location *LatLng `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *GeocodeResponse) Reset() {
	*x = GeocodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geocoder_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeocodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeocodeResponse) ProtoMessage() {}

func (x *GeocodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_geocoder_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeocodeResponse.ProtoReflect.Descriptor instead.
func (*GeocodeResponse) Descriptor() ([]byte, []int) {
	return file_geocoder_proto_rawDescGZIP(), []int{2}
}

func (x *GeocodeResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *GeocodeResponse) GetLocation() *LatLng {
	if x != nil {
		return x.Location
	}
	return nil
}

var File_geocoder_proto protoreflect.FileDescriptor

var file_geocoder_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x67, 0x65, 0x6f, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x67, 0x65, 0x6f, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x22, 0x42, 0x0a, 0x06, 0x4c, 0x61,
	0x74, 0x4c, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x3e,
	0x0a, 0x0e, 0x47, 0x65, 0x6f, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2c, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x65, 0x6f, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x2e, 0x4c, 0x61,
	0x74, 0x4c, 0x6e, 0x67, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x59,
	0x0a, 0x0f, 0x47, 0x65, 0x6f, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2c, 0x0a, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x67, 0x65, 0x6f, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x2e, 0x4c, 0x61, 0x74, 0x4c, 0x6e, 0x67, 0x52,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x53, 0x0a, 0x08, 0x47, 0x65, 0x6f,
	0x63, 0x6f, 0x64, 0x65, 0x72, 0x12, 0x47, 0x0a, 0x0e, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65,
	0x47, 0x65, 0x6f, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x2e, 0x67, 0x65, 0x6f, 0x63, 0x6f, 0x64,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x6f, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x67, 0x65, 0x6f, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x6f,
	0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_geocoder_proto_rawDescOnce sync.Once
	file_geocoder_proto_rawDescData = file_geocoder_proto_rawDesc
)

func file_geocoder_proto_rawDescGZIP() []byte {
	file_geocoder_proto_rawDescOnce.Do(func() {
		file_geocoder_proto_rawDescData = protoimpl.X.CompressGZIP(file_geocoder_proto_rawDescData)
	})
	return file_geocoder_proto_rawDescData
}

var file_geocoder_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_geocoder_proto_goTypes = []interface{}{
	(*LatLng)(nil),          // 0: geocoder.LatLng
	(*GeoCodeRequest)(nil),  // 1: geocoder.GeoCodeRequest
	(*GeocodeResponse)(nil), // 2: geocoder.GeocodeResponse
}
var file_geocoder_proto_depIdxs = []int32{
	0, // 0: geocoder.GeoCodeRequest.location:type_name -> geocoder.LatLng
	0, // 1: geocoder.GeocodeResponse.location:type_name -> geocoder.LatLng
	1, // 2: geocoder.Geocoder.ReverseGeoCode:input_type -> geocoder.GeoCodeRequest
	2, // 3: geocoder.Geocoder.ReverseGeoCode:output_type -> geocoder.GeocodeResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_geocoder_proto_init() }
func file_geocoder_proto_init() {
	if File_geocoder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_geocoder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LatLng); i {
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
		file_geocoder_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeoCodeRequest); i {
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
		file_geocoder_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeocodeResponse); i {
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
			RawDescriptor: file_geocoder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_geocoder_proto_goTypes,
		DependencyIndexes: file_geocoder_proto_depIdxs,
		MessageInfos:      file_geocoder_proto_msgTypes,
	}.Build()
	File_geocoder_proto = out.File
	file_geocoder_proto_rawDesc = nil
	file_geocoder_proto_goTypes = nil
	file_geocoder_proto_depIdxs = nil
}