// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: radish/weather/v1/weather.proto

package weatherv1

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

type Condition int32

const (
	Condition_CONDITION_UNSPECIFIED Condition = 0
	Condition_CONDITION_SUNNY       Condition = 1
	Condition_CONDITION_RAINY       Condition = 2
)

// Enum value maps for Condition.
var (
	Condition_name = map[int32]string{
		0: "CONDITION_UNSPECIFIED",
		1: "CONDITION_SUNNY",
		2: "CONDITION_RAINY",
	}
	Condition_value = map[string]int32{
		"CONDITION_UNSPECIFIED": 0,
		"CONDITION_SUNNY":       1,
		"CONDITION_RAINY":       2,
	}
)

func (x Condition) Enum() *Condition {
	p := new(Condition)
	*p = x
	return p
}

func (x Condition) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Condition) Descriptor() protoreflect.EnumDescriptor {
	return file_radish_weather_v1_weather_proto_enumTypes[0].Descriptor()
}

func (Condition) Type() protoreflect.EnumType {
	return &file_radish_weather_v1_weather_proto_enumTypes[0]
}

func (x Condition) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Condition.Descriptor instead.
func (Condition) EnumDescriptor() ([]byte, []int) {
	return file_radish_weather_v1_weather_proto_rawDescGZIP(), []int{0}
}

type GetWeatherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float32 `protobuf:"fixed32,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float32 `protobuf:"fixed32,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *GetWeatherRequest) Reset() {
	*x = GetWeatherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_radish_weather_v1_weather_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWeatherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWeatherRequest) ProtoMessage() {}

func (x *GetWeatherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_radish_weather_v1_weather_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWeatherRequest.ProtoReflect.Descriptor instead.
func (*GetWeatherRequest) Descriptor() ([]byte, []int) {
	return file_radish_weather_v1_weather_proto_rawDescGZIP(), []int{0}
}

func (x *GetWeatherRequest) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *GetWeatherRequest) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type GetWeatherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Temperature float32   `protobuf:"fixed32,1,opt,name=temperature,proto3" json:"temperature,omitempty"`
	Conditions  Condition `protobuf:"varint,2,opt,name=conditions,proto3,enum=radish.weather.v1.Condition" json:"conditions,omitempty"`
}

func (x *GetWeatherResponse) Reset() {
	*x = GetWeatherResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_radish_weather_v1_weather_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWeatherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWeatherResponse) ProtoMessage() {}

func (x *GetWeatherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_radish_weather_v1_weather_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWeatherResponse.ProtoReflect.Descriptor instead.
func (*GetWeatherResponse) Descriptor() ([]byte, []int) {
	return file_radish_weather_v1_weather_proto_rawDescGZIP(), []int{1}
}

func (x *GetWeatherResponse) GetTemperature() float32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *GetWeatherResponse) GetConditions() Condition {
	if x != nil {
		return x.Conditions
	}
	return Condition_CONDITION_UNSPECIFIED
}

var File_radish_weather_v1_weather_proto protoreflect.FileDescriptor

var file_radish_weather_v1_weather_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x72, 0x61, 0x64, 0x69, 0x73, 0x68, 0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x72, 0x61, 0x64, 0x69, 0x73, 0x68, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x22, 0x4d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x22, 0x74, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b,
	0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x3c, 0x0a, 0x0a, 0x63,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1c, 0x2e, 0x72, 0x61, 0x64, 0x69, 0x73, 0x68, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2a, 0x50, 0x0a, 0x09, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53,
	0x55, 0x4e, 0x4e, 0x59, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x41, 0x49, 0x4e, 0x59, 0x10, 0x02, 0x32, 0x6b, 0x0a, 0x0e, 0x57,
	0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x12, 0x24, 0x2e, 0x72, 0x61,
	0x64, 0x69, 0x73, 0x68, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x72, 0x61, 0x64, 0x69, 0x73, 0x68, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xdb, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d,
	0x2e, 0x72, 0x61, 0x64, 0x69, 0x73, 0x68, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x42, 0x0c, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72,
	0x61, 0x64, 0x69, 0x73, 0x68, 0x2d, 0x6d, 0x69, 0x79, 0x61, 0x7a, 0x61, 0x6b, 0x69, 0x2f, 0x62,
	0x75, 0x66, 0x2d, 0x76, 0x32, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x72, 0x61, 0x64, 0x69, 0x73, 0x68, 0x2f, 0x77, 0x65,
	0x61, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x52, 0x57, 0x58, 0xaa, 0x02, 0x11, 0x52, 0x61, 0x64, 0x69, 0x73,
	0x68, 0x2e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x11, 0x52,
	0x61, 0x64, 0x69, 0x73, 0x68, 0x5c, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x1d, 0x52, 0x61, 0x64, 0x69, 0x73, 0x68, 0x5c, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65,
	0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x13, 0x52, 0x61, 0x64, 0x69, 0x73, 0x68, 0x3a, 0x3a, 0x57, 0x65, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_radish_weather_v1_weather_proto_rawDescOnce sync.Once
	file_radish_weather_v1_weather_proto_rawDescData = file_radish_weather_v1_weather_proto_rawDesc
)

func file_radish_weather_v1_weather_proto_rawDescGZIP() []byte {
	file_radish_weather_v1_weather_proto_rawDescOnce.Do(func() {
		file_radish_weather_v1_weather_proto_rawDescData = protoimpl.X.CompressGZIP(file_radish_weather_v1_weather_proto_rawDescData)
	})
	return file_radish_weather_v1_weather_proto_rawDescData
}

var file_radish_weather_v1_weather_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_radish_weather_v1_weather_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_radish_weather_v1_weather_proto_goTypes = []any{
	(Condition)(0),             // 0: radish.weather.v1.Condition
	(*GetWeatherRequest)(nil),  // 1: radish.weather.v1.GetWeatherRequest
	(*GetWeatherResponse)(nil), // 2: radish.weather.v1.GetWeatherResponse
}
var file_radish_weather_v1_weather_proto_depIdxs = []int32{
	0, // 0: radish.weather.v1.GetWeatherResponse.conditions:type_name -> radish.weather.v1.Condition
	1, // 1: radish.weather.v1.WeatherService.GetWeather:input_type -> radish.weather.v1.GetWeatherRequest
	2, // 2: radish.weather.v1.WeatherService.GetWeather:output_type -> radish.weather.v1.GetWeatherResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_radish_weather_v1_weather_proto_init() }
func file_radish_weather_v1_weather_proto_init() {
	if File_radish_weather_v1_weather_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_radish_weather_v1_weather_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetWeatherRequest); i {
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
		file_radish_weather_v1_weather_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetWeatherResponse); i {
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
			RawDescriptor: file_radish_weather_v1_weather_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_radish_weather_v1_weather_proto_goTypes,
		DependencyIndexes: file_radish_weather_v1_weather_proto_depIdxs,
		EnumInfos:         file_radish_weather_v1_weather_proto_enumTypes,
		MessageInfos:      file_radish_weather_v1_weather_proto_msgTypes,
	}.Build()
	File_radish_weather_v1_weather_proto = out.File
	file_radish_weather_v1_weather_proto_rawDesc = nil
	file_radish_weather_v1_weather_proto_goTypes = nil
	file_radish_weather_v1_weather_proto_depIdxs = nil
}
