// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: flyt/monolith-order-created.proto

package flyt

import (
	ordering "github.com/flypay/events/pkg/monolith/ordering"
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

type MonolithOrderCreated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated: Do not use.
	Order                          *ordering.DeliveryByDeliveryPartnerOrder `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	DeliveryByOperatorOrder        *ordering.DeliveryByOperatorOrder        `protobuf:"bytes,2,opt,name=delivery_by_operator_order,json=deliveryByOperatorOrder,proto3" json:"delivery_by_operator_order,omitempty"`
	CollectionByCustomerOrder      *ordering.CollectionOrder                `protobuf:"bytes,3,opt,name=collection_by_customer_order,json=collectionByCustomerOrder,proto3" json:"collection_by_customer_order,omitempty"`
	Restaurant                     *RestaurantIdent                         `protobuf:"bytes,4,opt,name=restaurant,proto3" json:"restaurant,omitempty"`
	DeliveryByDeliveryPartnerOrder *ordering.DeliveryByDeliveryPartnerOrder `protobuf:"bytes,5,opt,name=delivery_by_delivery_partner_order,json=deliveryByDeliveryPartnerOrder,proto3" json:"delivery_by_delivery_partner_order,omitempty"`
	TestData                       *TestData                                `protobuf:"bytes,16,opt,name=test_data,json=testData,proto3" json:"test_data,omitempty"`
	TransmissionId                 string                                   `protobuf:"bytes,17,opt,name=transmission_id,json=transmissionId,proto3" json:"transmission_id,omitempty"`
}

func (x *MonolithOrderCreated) Reset() {
	*x = MonolithOrderCreated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyt_monolith_order_created_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonolithOrderCreated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonolithOrderCreated) ProtoMessage() {}

func (x *MonolithOrderCreated) ProtoReflect() protoreflect.Message {
	mi := &file_flyt_monolith_order_created_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonolithOrderCreated.ProtoReflect.Descriptor instead.
func (*MonolithOrderCreated) Descriptor() ([]byte, []int) {
	return file_flyt_monolith_order_created_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Do not use.
func (x *MonolithOrderCreated) GetOrder() *ordering.DeliveryByDeliveryPartnerOrder {
	if x != nil {
		return x.Order
	}
	return nil
}

func (x *MonolithOrderCreated) GetDeliveryByOperatorOrder() *ordering.DeliveryByOperatorOrder {
	if x != nil {
		return x.DeliveryByOperatorOrder
	}
	return nil
}

func (x *MonolithOrderCreated) GetCollectionByCustomerOrder() *ordering.CollectionOrder {
	if x != nil {
		return x.CollectionByCustomerOrder
	}
	return nil
}

func (x *MonolithOrderCreated) GetRestaurant() *RestaurantIdent {
	if x != nil {
		return x.Restaurant
	}
	return nil
}

func (x *MonolithOrderCreated) GetDeliveryByDeliveryPartnerOrder() *ordering.DeliveryByDeliveryPartnerOrder {
	if x != nil {
		return x.DeliveryByDeliveryPartnerOrder
	}
	return nil
}

func (x *MonolithOrderCreated) GetTestData() *TestData {
	if x != nil {
		return x.TestData
	}
	return nil
}

func (x *MonolithOrderCreated) GetTransmissionId() string {
	if x != nil {
		return x.TransmissionId
	}
	return ""
}

var File_flyt_monolith_order_created_proto protoreflect.FileDescriptor

var file_flyt_monolith_order_created_proto_rawDesc = []byte{
	0x0a, 0x21, 0x66, 0x6c, 0x79, 0x74, 0x2f, 0x6d, 0x6f, 0x6e, 0x6f, 0x6c, 0x69, 0x74, 0x68, 0x2d,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x6c, 0x79, 0x74, 0x1a, 0x15, 0x66, 0x6c, 0x79, 0x74, 0x2f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x66, 0x6c, 0x79, 0x74, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x66, 0x6c, 0x79, 0x74, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6d,
	0x6f, 0x6e, 0x6f, 0x6c, 0x69, 0x74, 0x68, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67,
	0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb5, 0x04, 0x0a, 0x14, 0x4d, 0x6f, 0x6e, 0x6f, 0x6c, 0x69, 0x74, 0x68, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x42, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69,
	0x6e, 0x67, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x42, 0x79, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x42, 0x02, 0x18, 0x01, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x5e, 0x0a, 0x1a,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x62, 0x79, 0x5f, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x42, 0x79, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x17, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x42, 0x79, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x5a, 0x0a, 0x1c,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x79, 0x5f, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x19, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x74,
	0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x66,
	0x6c, 0x79, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x12,
	0x74, 0x0a, 0x22, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x62, 0x79, 0x5f, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x42,
	0x79, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x1e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x42,
	0x79, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x09, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x2e,
	0x54, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x74, 0x65, 0x73, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x3a, 0x1a, 0x82, 0xb5, 0x18,
	0x16, 0x6d, 0x6f, 0x6e, 0x6f, 0x6c, 0x69, 0x74, 0x68, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x78, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e, 0x66,
	0x6c, 0x79, 0x74, 0x42, 0x19, 0x4d, 0x6f, 0x6e, 0x6f, 0x6c, 0x69, 0x74, 0x68, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x79,
	0x70, 0x61, 0x79, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x66,
	0x6c, 0x79, 0x74, 0xa2, 0x02, 0x03, 0x46, 0x58, 0x58, 0xaa, 0x02, 0x04, 0x46, 0x6c, 0x79, 0x74,
	0xca, 0x02, 0x04, 0x46, 0x6c, 0x79, 0x74, 0xe2, 0x02, 0x10, 0x46, 0x6c, 0x79, 0x74, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x04, 0x46, 0x6c, 0x79,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flyt_monolith_order_created_proto_rawDescOnce sync.Once
	file_flyt_monolith_order_created_proto_rawDescData = file_flyt_monolith_order_created_proto_rawDesc
)

func file_flyt_monolith_order_created_proto_rawDescGZIP() []byte {
	file_flyt_monolith_order_created_proto_rawDescOnce.Do(func() {
		file_flyt_monolith_order_created_proto_rawDescData = protoimpl.X.CompressGZIP(file_flyt_monolith_order_created_proto_rawDescData)
	})
	return file_flyt_monolith_order_created_proto_rawDescData
}

var file_flyt_monolith_order_created_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_flyt_monolith_order_created_proto_goTypes = []interface{}{
	(*MonolithOrderCreated)(nil),                    // 0: flyt.MonolithOrderCreated
	(*ordering.DeliveryByDeliveryPartnerOrder)(nil), // 1: ordering.DeliveryByDeliveryPartnerOrder
	(*ordering.DeliveryByOperatorOrder)(nil),        // 2: ordering.DeliveryByOperatorOrder
	(*ordering.CollectionOrder)(nil),                // 3: ordering.CollectionOrder
	(*RestaurantIdent)(nil),                         // 4: flyt.RestaurantIdent
	(*TestData)(nil),                                // 5: flyt.TestData
}
var file_flyt_monolith_order_created_proto_depIdxs = []int32{
	1, // 0: flyt.MonolithOrderCreated.order:type_name -> ordering.DeliveryByDeliveryPartnerOrder
	2, // 1: flyt.MonolithOrderCreated.delivery_by_operator_order:type_name -> ordering.DeliveryByOperatorOrder
	3, // 2: flyt.MonolithOrderCreated.collection_by_customer_order:type_name -> ordering.CollectionOrder
	4, // 3: flyt.MonolithOrderCreated.restaurant:type_name -> flyt.RestaurantIdent
	1, // 4: flyt.MonolithOrderCreated.delivery_by_delivery_partner_order:type_name -> ordering.DeliveryByDeliveryPartnerOrder
	5, // 5: flyt.MonolithOrderCreated.test_data:type_name -> flyt.TestData
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_flyt_monolith_order_created_proto_init() }
func file_flyt_monolith_order_created_proto_init() {
	if File_flyt_monolith_order_created_proto != nil {
		return
	}
	file_flyt_descriptor_proto_init()
	file_flyt_restaurant_proto_init()
	file_flyt_test_data_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_flyt_monolith_order_created_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonolithOrderCreated); i {
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
			RawDescriptor: file_flyt_monolith_order_created_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flyt_monolith_order_created_proto_goTypes,
		DependencyIndexes: file_flyt_monolith_order_created_proto_depIdxs,
		MessageInfos:      file_flyt_monolith_order_created_proto_msgTypes,
	}.Build()
	File_flyt_monolith_order_created_proto = out.File
	file_flyt_monolith_order_created_proto_rawDesc = nil
	file_flyt_monolith_order_created_proto_goTypes = nil
	file_flyt_monolith_order_created_proto_depIdxs = nil
}
