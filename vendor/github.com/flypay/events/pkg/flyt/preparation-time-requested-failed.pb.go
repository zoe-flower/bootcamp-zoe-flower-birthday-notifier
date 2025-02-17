// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: flyt/preparation-time-requested-failed.proto

package flyt

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

type PreparationTimeRequestedFailed_ErrorCode int32

const (
	PreparationTimeRequestedFailed_UNKNOWN            PreparationTimeRequestedFailed_ErrorCode = 0
	PreparationTimeRequestedFailed_BAD_REQUEST        PreparationTimeRequestedFailed_ErrorCode = 1
	PreparationTimeRequestedFailed_AUTH_FAILED        PreparationTimeRequestedFailed_ErrorCode = 2
	PreparationTimeRequestedFailed_CONNECTIVITY_ISSUE PreparationTimeRequestedFailed_ErrorCode = 3
	PreparationTimeRequestedFailed_TIMEOUT            PreparationTimeRequestedFailed_ErrorCode = 4
)

// Enum value maps for PreparationTimeRequestedFailed_ErrorCode.
var (
	PreparationTimeRequestedFailed_ErrorCode_name = map[int32]string{
		0: "UNKNOWN",
		1: "BAD_REQUEST",
		2: "AUTH_FAILED",
		3: "CONNECTIVITY_ISSUE",
		4: "TIMEOUT",
	}
	PreparationTimeRequestedFailed_ErrorCode_value = map[string]int32{
		"UNKNOWN":            0,
		"BAD_REQUEST":        1,
		"AUTH_FAILED":        2,
		"CONNECTIVITY_ISSUE": 3,
		"TIMEOUT":            4,
	}
)

func (x PreparationTimeRequestedFailed_ErrorCode) Enum() *PreparationTimeRequestedFailed_ErrorCode {
	p := new(PreparationTimeRequestedFailed_ErrorCode)
	*p = x
	return p
}

func (x PreparationTimeRequestedFailed_ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PreparationTimeRequestedFailed_ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_flyt_preparation_time_requested_failed_proto_enumTypes[0].Descriptor()
}

func (PreparationTimeRequestedFailed_ErrorCode) Type() protoreflect.EnumType {
	return &file_flyt_preparation_time_requested_failed_proto_enumTypes[0]
}

func (x PreparationTimeRequestedFailed_ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PreparationTimeRequestedFailed_ErrorCode.Descriptor instead.
func (PreparationTimeRequestedFailed_ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_flyt_preparation_time_requested_failed_proto_rawDescGZIP(), []int{0, 0}
}

type PreparationTimeRequestedFailed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Define your proto fields here
	Version int32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Mark as deprecated. Wrong ident type.
	//
	// Deprecated: Do not use.
	Restaurant *RestaurantIdent                      `protobuf:"bytes,2,opt,name=restaurant,proto3" json:"restaurant,omitempty"`
	Error      *PreparationTimeRequestedFailed_Error `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	// Add correct restaurant identifier field
	RestaurantId *Ident `protobuf:"bytes,4,opt,name=restaurant_id,json=restaurantId,proto3" json:"restaurant_id,omitempty"`
	// Add external order ID reference
	ExternalReference string `protobuf:"bytes,5,opt,name=external_reference,json=externalReference,proto3" json:"external_reference,omitempty"`
}

func (x *PreparationTimeRequestedFailed) Reset() {
	*x = PreparationTimeRequestedFailed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyt_preparation_time_requested_failed_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreparationTimeRequestedFailed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreparationTimeRequestedFailed) ProtoMessage() {}

func (x *PreparationTimeRequestedFailed) ProtoReflect() protoreflect.Message {
	mi := &file_flyt_preparation_time_requested_failed_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreparationTimeRequestedFailed.ProtoReflect.Descriptor instead.
func (*PreparationTimeRequestedFailed) Descriptor() ([]byte, []int) {
	return file_flyt_preparation_time_requested_failed_proto_rawDescGZIP(), []int{0}
}

func (x *PreparationTimeRequestedFailed) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

// Deprecated: Do not use.
func (x *PreparationTimeRequestedFailed) GetRestaurant() *RestaurantIdent {
	if x != nil {
		return x.Restaurant
	}
	return nil
}

func (x *PreparationTimeRequestedFailed) GetError() *PreparationTimeRequestedFailed_Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *PreparationTimeRequestedFailed) GetRestaurantId() *Ident {
	if x != nil {
		return x.RestaurantId
	}
	return nil
}

func (x *PreparationTimeRequestedFailed) GetExternalReference() string {
	if x != nil {
		return x.ExternalReference
	}
	return ""
}

type PreparationTimeRequestedFailed_Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    PreparationTimeRequestedFailed_ErrorCode `protobuf:"varint,1,opt,name=code,proto3,enum=flyt.PreparationTimeRequestedFailed_ErrorCode" json:"code,omitempty"`
	Message string                                   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PreparationTimeRequestedFailed_Error) Reset() {
	*x = PreparationTimeRequestedFailed_Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyt_preparation_time_requested_failed_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreparationTimeRequestedFailed_Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreparationTimeRequestedFailed_Error) ProtoMessage() {}

func (x *PreparationTimeRequestedFailed_Error) ProtoReflect() protoreflect.Message {
	mi := &file_flyt_preparation_time_requested_failed_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreparationTimeRequestedFailed_Error.ProtoReflect.Descriptor instead.
func (*PreparationTimeRequestedFailed_Error) Descriptor() ([]byte, []int) {
	return file_flyt_preparation_time_requested_failed_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PreparationTimeRequestedFailed_Error) GetCode() PreparationTimeRequestedFailed_ErrorCode {
	if x != nil {
		return x.Code
	}
	return PreparationTimeRequestedFailed_UNKNOWN
}

func (x *PreparationTimeRequestedFailed_Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_flyt_preparation_time_requested_failed_proto protoreflect.FileDescriptor

var file_flyt_preparation_time_requested_failed_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x66, 0x6c, 0x79, 0x74, 0x2f, 0x70, 0x72, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2d, 0x74, 0x69, 0x6d, 0x65, 0x2d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x2d, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x66, 0x6c, 0x79, 0x74, 0x1a, 0x15, 0x66, 0x6c, 0x79, 0x74, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x66, 0x6c, 0x79,
	0x74, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x66,
	0x6c, 0x79, 0x74, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x04, 0x0a, 0x1e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x2e, 0x52, 0x65, 0x73,
	0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x42, 0x02, 0x18, 0x01,
	0x52, 0x0a, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x40, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x66, 0x6c,
	0x79, 0x74, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x46, 0x61, 0x69, 0x6c, 0x65,
	0x64, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x30,
	0x0a, 0x0d, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x2e, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x2d, 0x0a, 0x12, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x1a,
	0x65, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x42, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x2e, 0x50, 0x72,
	0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x5f, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x41, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10,
	0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44,
	0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x56, 0x49,
	0x54, 0x59, 0x5f, 0x49, 0x53, 0x53, 0x55, 0x45, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x49,
	0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x04, 0x3a, 0x25, 0x82, 0xb5, 0x18, 0x21, 0x70, 0x72, 0x65,
	0x70, 0x61, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x2e, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x42, 0x82,
	0x01, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x42, 0x23, 0x50, 0x72, 0x65,
	0x70, 0x61, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66,
	0x6c, 0x79, 0x70, 0x61, 0x79, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x66, 0x6c, 0x79, 0x74, 0xa2, 0x02, 0x03, 0x46, 0x58, 0x58, 0xaa, 0x02, 0x04, 0x46, 0x6c,
	0x79, 0x74, 0xca, 0x02, 0x04, 0x46, 0x6c, 0x79, 0x74, 0xe2, 0x02, 0x10, 0x46, 0x6c, 0x79, 0x74,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x04, 0x46,
	0x6c, 0x79, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flyt_preparation_time_requested_failed_proto_rawDescOnce sync.Once
	file_flyt_preparation_time_requested_failed_proto_rawDescData = file_flyt_preparation_time_requested_failed_proto_rawDesc
)

func file_flyt_preparation_time_requested_failed_proto_rawDescGZIP() []byte {
	file_flyt_preparation_time_requested_failed_proto_rawDescOnce.Do(func() {
		file_flyt_preparation_time_requested_failed_proto_rawDescData = protoimpl.X.CompressGZIP(file_flyt_preparation_time_requested_failed_proto_rawDescData)
	})
	return file_flyt_preparation_time_requested_failed_proto_rawDescData
}

var file_flyt_preparation_time_requested_failed_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_flyt_preparation_time_requested_failed_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_flyt_preparation_time_requested_failed_proto_goTypes = []interface{}{
	(PreparationTimeRequestedFailed_ErrorCode)(0), // 0: flyt.PreparationTimeRequestedFailed.ErrorCode
	(*PreparationTimeRequestedFailed)(nil),        // 1: flyt.PreparationTimeRequestedFailed
	(*PreparationTimeRequestedFailed_Error)(nil),  // 2: flyt.PreparationTimeRequestedFailed.Error
	(*RestaurantIdent)(nil),                       // 3: flyt.RestaurantIdent
	(*Ident)(nil),                                 // 4: flyt.Ident
}
var file_flyt_preparation_time_requested_failed_proto_depIdxs = []int32{
	3, // 0: flyt.PreparationTimeRequestedFailed.restaurant:type_name -> flyt.RestaurantIdent
	2, // 1: flyt.PreparationTimeRequestedFailed.error:type_name -> flyt.PreparationTimeRequestedFailed.Error
	4, // 2: flyt.PreparationTimeRequestedFailed.restaurant_id:type_name -> flyt.Ident
	0, // 3: flyt.PreparationTimeRequestedFailed.Error.code:type_name -> flyt.PreparationTimeRequestedFailed.ErrorCode
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_flyt_preparation_time_requested_failed_proto_init() }
func file_flyt_preparation_time_requested_failed_proto_init() {
	if File_flyt_preparation_time_requested_failed_proto != nil {
		return
	}
	file_flyt_descriptor_proto_init()
	file_flyt_ident_proto_init()
	file_flyt_restaurant_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_flyt_preparation_time_requested_failed_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreparationTimeRequestedFailed); i {
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
		file_flyt_preparation_time_requested_failed_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreparationTimeRequestedFailed_Error); i {
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
			RawDescriptor: file_flyt_preparation_time_requested_failed_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flyt_preparation_time_requested_failed_proto_goTypes,
		DependencyIndexes: file_flyt_preparation_time_requested_failed_proto_depIdxs,
		EnumInfos:         file_flyt_preparation_time_requested_failed_proto_enumTypes,
		MessageInfos:      file_flyt_preparation_time_requested_failed_proto_msgTypes,
	}.Build()
	File_flyt_preparation_time_requested_failed_proto = out.File
	file_flyt_preparation_time_requested_failed_proto_rawDesc = nil
	file_flyt_preparation_time_requested_failed_proto_goTypes = nil
	file_flyt_preparation_time_requested_failed_proto_depIdxs = nil
}
