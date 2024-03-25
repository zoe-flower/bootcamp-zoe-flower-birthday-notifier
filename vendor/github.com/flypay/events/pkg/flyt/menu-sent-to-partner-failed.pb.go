// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: flyt/menu-sent-to-partner-failed.proto

package flyt

import (
	menutype "github.com/flypay/events/pkg/menutype"
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

// When a menu failed to be sent to partner
type MenuSentToPartnerFailed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Restaurant      *Ident             `protobuf:"bytes,1,opt,name=restaurant,proto3" json:"restaurant,omitempty"`
	Type            menutype.Menu_Type `protobuf:"varint,2,opt,name=type,proto3,enum=menutype.Menu_Type" json:"type,omitempty"`
	Error           string             `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"` // error message
	ErrorCode       MenuErrorCode      `protobuf:"varint,4,opt,name=error_code,json=errorCode,proto3,enum=flyt.MenuErrorCode" json:"error_code,omitempty"`
	ErrorDetails    string             `protobuf:"bytes,5,opt,name=error_details,json=errorDetails,proto3" json:"error_details,omitempty"`
	ErrorResolution string             `protobuf:"bytes,6,opt,name=error_resolution,json=errorResolution,proto3" json:"error_resolution,omitempty"`
	// Unique identifier, originating from the published menus
	MenuId string `protobuf:"bytes,7,opt,name=menu_id,json=menuId,proto3" json:"menu_id,omitempty"`
}

func (x *MenuSentToPartnerFailed) Reset() {
	*x = MenuSentToPartnerFailed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyt_menu_sent_to_partner_failed_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuSentToPartnerFailed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuSentToPartnerFailed) ProtoMessage() {}

func (x *MenuSentToPartnerFailed) ProtoReflect() protoreflect.Message {
	mi := &file_flyt_menu_sent_to_partner_failed_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuSentToPartnerFailed.ProtoReflect.Descriptor instead.
func (*MenuSentToPartnerFailed) Descriptor() ([]byte, []int) {
	return file_flyt_menu_sent_to_partner_failed_proto_rawDescGZIP(), []int{0}
}

func (x *MenuSentToPartnerFailed) GetRestaurant() *Ident {
	if x != nil {
		return x.Restaurant
	}
	return nil
}

func (x *MenuSentToPartnerFailed) GetType() menutype.Menu_Type {
	if x != nil {
		return x.Type
	}
	return menutype.Menu_ANY
}

func (x *MenuSentToPartnerFailed) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *MenuSentToPartnerFailed) GetErrorCode() MenuErrorCode {
	if x != nil {
		return x.ErrorCode
	}
	return MenuErrorCode_ERROR_UNKNOWN
}

func (x *MenuSentToPartnerFailed) GetErrorDetails() string {
	if x != nil {
		return x.ErrorDetails
	}
	return ""
}

func (x *MenuSentToPartnerFailed) GetErrorResolution() string {
	if x != nil {
		return x.ErrorResolution
	}
	return ""
}

func (x *MenuSentToPartnerFailed) GetMenuId() string {
	if x != nil {
		return x.MenuId
	}
	return ""
}

var File_flyt_menu_sent_to_partner_failed_proto protoreflect.FileDescriptor

var file_flyt_menu_sent_to_partner_failed_proto_rawDesc = []byte{
	0x0a, 0x26, 0x66, 0x6c, 0x79, 0x74, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2d, 0x73, 0x65, 0x6e, 0x74,
	0x2d, 0x74, 0x6f, 0x2d, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2d, 0x66, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x6c, 0x79, 0x74, 0x1a, 0x15,
	0x66, 0x6c, 0x79, 0x74, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x66, 0x6c, 0x79, 0x74, 0x2f, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x66, 0x6c, 0x79, 0x74, 0x2f, 0x6d, 0x65,
	0x6e, 0x75, 0x2d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2d, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x6d, 0x65, 0x6e, 0x75, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x65,
	0x6e, 0x75, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x02, 0x0a,
	0x17, 0x4d, 0x65, 0x6e, 0x75, 0x53, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x50, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x2b, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x74,
	0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x66,
	0x6c, 0x79, 0x74, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x74, 0x61,
	0x75, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x4d,
	0x65, 0x6e, 0x75, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x32, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x2e,
	0x4d, 0x65, 0x6e, 0x75, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x29, 0x0a,
	0x10, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65,
	0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x65, 0x6e, 0x75,
	0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x6e, 0x75, 0x49,
	0x64, 0x3a, 0x1f, 0x82, 0xb5, 0x18, 0x1b, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x73, 0x65, 0x6e, 0x74,
	0x2e, 0x74, 0x6f, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x66, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x42, 0x7b, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x42, 0x1c,
	0x4d, 0x65, 0x6e, 0x75, 0x53, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x21,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x79, 0x70, 0x61,
	0x79, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x6c, 0x79,
	0x74, 0xa2, 0x02, 0x03, 0x46, 0x58, 0x58, 0xaa, 0x02, 0x04, 0x46, 0x6c, 0x79, 0x74, 0xca, 0x02,
	0x04, 0x46, 0x6c, 0x79, 0x74, 0xe2, 0x02, 0x10, 0x46, 0x6c, 0x79, 0x74, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x04, 0x46, 0x6c, 0x79, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flyt_menu_sent_to_partner_failed_proto_rawDescOnce sync.Once
	file_flyt_menu_sent_to_partner_failed_proto_rawDescData = file_flyt_menu_sent_to_partner_failed_proto_rawDesc
)

func file_flyt_menu_sent_to_partner_failed_proto_rawDescGZIP() []byte {
	file_flyt_menu_sent_to_partner_failed_proto_rawDescOnce.Do(func() {
		file_flyt_menu_sent_to_partner_failed_proto_rawDescData = protoimpl.X.CompressGZIP(file_flyt_menu_sent_to_partner_failed_proto_rawDescData)
	})
	return file_flyt_menu_sent_to_partner_failed_proto_rawDescData
}

var file_flyt_menu_sent_to_partner_failed_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_flyt_menu_sent_to_partner_failed_proto_goTypes = []interface{}{
	(*MenuSentToPartnerFailed)(nil), // 0: flyt.MenuSentToPartnerFailed
	(*Ident)(nil),                   // 1: flyt.Ident
	(menutype.Menu_Type)(0),         // 2: menutype.Menu.Type
	(MenuErrorCode)(0),              // 3: flyt.MenuErrorCode
}
var file_flyt_menu_sent_to_partner_failed_proto_depIdxs = []int32{
	1, // 0: flyt.MenuSentToPartnerFailed.restaurant:type_name -> flyt.Ident
	2, // 1: flyt.MenuSentToPartnerFailed.type:type_name -> menutype.Menu.Type
	3, // 2: flyt.MenuSentToPartnerFailed.error_code:type_name -> flyt.MenuErrorCode
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_flyt_menu_sent_to_partner_failed_proto_init() }
func file_flyt_menu_sent_to_partner_failed_proto_init() {
	if File_flyt_menu_sent_to_partner_failed_proto != nil {
		return
	}
	file_flyt_descriptor_proto_init()
	file_flyt_ident_proto_init()
	file_flyt_menu_error_code_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_flyt_menu_sent_to_partner_failed_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuSentToPartnerFailed); i {
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
			RawDescriptor: file_flyt_menu_sent_to_partner_failed_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flyt_menu_sent_to_partner_failed_proto_goTypes,
		DependencyIndexes: file_flyt_menu_sent_to_partner_failed_proto_depIdxs,
		MessageInfos:      file_flyt_menu_sent_to_partner_failed_proto_msgTypes,
	}.Build()
	File_flyt_menu_sent_to_partner_failed_proto = out.File
	file_flyt_menu_sent_to_partner_failed_proto_rawDesc = nil
	file_flyt_menu_sent_to_partner_failed_proto_goTypes = nil
	file_flyt_menu_sent_to_partner_failed_proto_depIdxs = nil
}
