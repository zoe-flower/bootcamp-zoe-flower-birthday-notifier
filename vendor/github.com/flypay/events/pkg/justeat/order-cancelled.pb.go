// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: justeat/order-cancelled.proto

package justeat

import (
	_ "github.com/flypay/events/pkg/flyt"
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

type OrderCancelled struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerAware         bool                              `protobuf:"varint,1,opt,name=customer_aware,json=CustomerAware,proto3" json:"customer_aware,omitempty"`
	CustomerReference     string                            `protobuf:"bytes,2,opt,name=customer_reference,json=CustomerReference,proto3" json:"customer_reference,omitempty"`
	CustomerTransaction   *OrderCancelled_Transaction       `protobuf:"bytes,3,opt,name=customer_transaction,json=CustomerTransaction,proto3" json:"customer_transaction,omitempty"`
	InitiatedBy           string                            `protobuf:"bytes,4,opt,name=initiated_by,json=InitiatedBy,proto3" json:"initiated_by,omitempty"`
	Notes                 string                            `protobuf:"bytes,5,opt,name=notes,json=Notes,proto3" json:"notes,omitempty"`
	OrderId               string                            `protobuf:"bytes,6,opt,name=order_id,json=OrderId,proto3" json:"order_id,omitempty"`
	Reason                *OrderCancelled_CancelationReason `protobuf:"bytes,7,opt,name=reason,json=Reason,proto3" json:"reason,omitempty"`
	RestaurantAware       bool                              `protobuf:"varint,8,opt,name=restaurant_aware,json=RestaurantAware,proto3" json:"restaurant_aware,omitempty"`
	RestaurantId          string                            `protobuf:"bytes,9,opt,name=restaurant_id,json=RestaurantId,proto3" json:"restaurant_id,omitempty"`
	RestaurantTransaction *OrderCancelled_Transaction       `protobuf:"bytes,10,opt,name=restaurant_transaction,json=RestaurantTransaction,proto3" json:"restaurant_transaction,omitempty"`
	UserName              string                            `protobuf:"bytes,11,opt,name=user_name,json=UserName,proto3" json:"user_name,omitempty"`
	TimeStamp             string                            `protobuf:"bytes,12,opt,name=time_stamp,json=TimeStamp,proto3" json:"time_stamp,omitempty"`
	RaisingComponent      string                            `protobuf:"bytes,13,opt,name=raising_component,json=RaisingComponent,proto3" json:"raising_component,omitempty"`
	Tenant                string                            `protobuf:"bytes,14,opt,name=tenant,json=Tenant,proto3" json:"tenant,omitempty"`
	Conversation          string                            `protobuf:"bytes,15,opt,name=conversation,json=Conversation,proto3" json:"conversation,omitempty"`
}

func (x *OrderCancelled) Reset() {
	*x = OrderCancelled{}
	if protoimpl.UnsafeEnabled {
		mi := &file_justeat_order_cancelled_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCancelled) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCancelled) ProtoMessage() {}

func (x *OrderCancelled) ProtoReflect() protoreflect.Message {
	mi := &file_justeat_order_cancelled_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCancelled.ProtoReflect.Descriptor instead.
func (*OrderCancelled) Descriptor() ([]byte, []int) {
	return file_justeat_order_cancelled_proto_rawDescGZIP(), []int{0}
}

func (x *OrderCancelled) GetCustomerAware() bool {
	if x != nil {
		return x.CustomerAware
	}
	return false
}

func (x *OrderCancelled) GetCustomerReference() string {
	if x != nil {
		return x.CustomerReference
	}
	return ""
}

func (x *OrderCancelled) GetCustomerTransaction() *OrderCancelled_Transaction {
	if x != nil {
		return x.CustomerTransaction
	}
	return nil
}

func (x *OrderCancelled) GetInitiatedBy() string {
	if x != nil {
		return x.InitiatedBy
	}
	return ""
}

func (x *OrderCancelled) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

func (x *OrderCancelled) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *OrderCancelled) GetReason() *OrderCancelled_CancelationReason {
	if x != nil {
		return x.Reason
	}
	return nil
}

func (x *OrderCancelled) GetRestaurantAware() bool {
	if x != nil {
		return x.RestaurantAware
	}
	return false
}

func (x *OrderCancelled) GetRestaurantId() string {
	if x != nil {
		return x.RestaurantId
	}
	return ""
}

func (x *OrderCancelled) GetRestaurantTransaction() *OrderCancelled_Transaction {
	if x != nil {
		return x.RestaurantTransaction
	}
	return nil
}

func (x *OrderCancelled) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *OrderCancelled) GetTimeStamp() string {
	if x != nil {
		return x.TimeStamp
	}
	return ""
}

func (x *OrderCancelled) GetRaisingComponent() string {
	if x != nil {
		return x.RaisingComponent
	}
	return ""
}

func (x *OrderCancelled) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *OrderCancelled) GetConversation() string {
	if x != nil {
		return x.Conversation
	}
	return ""
}

type OrderCancelled_Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount     float64 `protobuf:"fixed64,1,opt,name=amount,json=Amount,proto3" json:"amount,omitempty"`
	Percentage float64 `protobuf:"fixed64,2,opt,name=percentage,json=Percentage,proto3" json:"percentage,omitempty"`
	Type       string  `protobuf:"bytes,3,opt,name=type,json=Type,proto3" json:"type,omitempty"`
}

func (x *OrderCancelled_Transaction) Reset() {
	*x = OrderCancelled_Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_justeat_order_cancelled_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCancelled_Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCancelled_Transaction) ProtoMessage() {}

func (x *OrderCancelled_Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_justeat_order_cancelled_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCancelled_Transaction.ProtoReflect.Descriptor instead.
func (*OrderCancelled_Transaction) Descriptor() ([]byte, []int) {
	return file_justeat_order_cancelled_proto_rawDescGZIP(), []int{0, 0}
}

func (x *OrderCancelled_Transaction) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *OrderCancelled_Transaction) GetPercentage() float64 {
	if x != nil {
		return x.Percentage
	}
	return 0
}

func (x *OrderCancelled_Transaction) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type OrderCancelled_CancelationReason struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string `protobuf:"bytes,1,opt,name=code,json=Code,proto3" json:"code,omitempty"`
	Comment string `protobuf:"bytes,2,opt,name=comment,json=Comment,proto3" json:"comment,omitempty"`
	SubCode string `protobuf:"bytes,3,opt,name=sub_code,json=SubCode,proto3" json:"sub_code,omitempty"`
}

func (x *OrderCancelled_CancelationReason) Reset() {
	*x = OrderCancelled_CancelationReason{}
	if protoimpl.UnsafeEnabled {
		mi := &file_justeat_order_cancelled_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCancelled_CancelationReason) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCancelled_CancelationReason) ProtoMessage() {}

func (x *OrderCancelled_CancelationReason) ProtoReflect() protoreflect.Message {
	mi := &file_justeat_order_cancelled_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCancelled_CancelationReason.ProtoReflect.Descriptor instead.
func (*OrderCancelled_CancelationReason) Descriptor() ([]byte, []int) {
	return file_justeat_order_cancelled_proto_rawDescGZIP(), []int{0, 1}
}

func (x *OrderCancelled_CancelationReason) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *OrderCancelled_CancelationReason) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *OrderCancelled_CancelationReason) GetSubCode() string {
	if x != nil {
		return x.SubCode
	}
	return ""
}

var File_justeat_order_cancelled_proto protoreflect.FileDescriptor

var file_justeat_order_cancelled_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2d,
	0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0x1a, 0x15, 0x66, 0x6c, 0x79, 0x74, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x15, 0x6a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x07, 0x0a, 0x0e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x61, 0x77, 0x61, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0d, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x41, 0x77, 0x61, 0x72, 0x65,
	0x12, 0x2d, 0x0a, 0x12, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12,
	0x56, 0x0a, 0x14, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x6a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x13, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x49,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f,
	0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4e, 0x6f, 0x74, 0x65, 0x73,
	0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x06, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6a, 0x75,
	0x73, 0x74, 0x65, 0x61, 0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x6c, 0x65, 0x64, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x29,
	0x0a, 0x10, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x61, 0x77, 0x61,
	0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x41, 0x77, 0x61, 0x72, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73,
	0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x5a,
	0x0a, 0x16, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x6a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x15, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x69, 0x6d,
	0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2b, 0x0a, 0x11, 0x72, 0x61, 0x69, 0x73, 0x69, 0x6e,
	0x67, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x52, 0x61, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x63,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x59, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x50, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x5c, 0x0a, 0x11, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x73, 0x75, 0x62, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x53, 0x75, 0x62, 0x43, 0x6f, 0x64, 0x65, 0x3a, 0x51, 0x82, 0xb5, 0x18, 0x17, 0x6a, 0x75,
	0x73, 0x74, 0x65, 0x61, 0x74, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x63, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x6c, 0x65, 0x64, 0xa2, 0xbb, 0x18, 0x0e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0xaa, 0xbb, 0x18, 0x02, 0x75, 0x6b, 0xaa, 0xbb, 0x18,
	0x02, 0x65, 0x73, 0xaa, 0xbb, 0x18, 0x02, 0x69, 0x65, 0xaa, 0xbb, 0x18, 0x02, 0x69, 0x74, 0xaa,
	0xbb, 0x18, 0x02, 0x61, 0x75, 0xaa, 0xbb, 0x18, 0x02, 0x6e, 0x7a, 0x42, 0x84, 0x01, 0x0a, 0x0b,
	0x63, 0x6f, 0x6d, 0x2e, 0x6a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0x42, 0x13, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66,
	0x6c, 0x79, 0x70, 0x61, 0x79, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x6a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0xa2, 0x02, 0x03, 0x4a, 0x58, 0x58, 0xaa, 0x02,
	0x07, 0x4a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0xca, 0x02, 0x07, 0x4a, 0x75, 0x73, 0x74, 0x65,
	0x61, 0x74, 0xe2, 0x02, 0x13, 0x4a, 0x75, 0x73, 0x74, 0x65, 0x61, 0x74, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x4a, 0x75, 0x73, 0x74, 0x65,
	0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_justeat_order_cancelled_proto_rawDescOnce sync.Once
	file_justeat_order_cancelled_proto_rawDescData = file_justeat_order_cancelled_proto_rawDesc
)

func file_justeat_order_cancelled_proto_rawDescGZIP() []byte {
	file_justeat_order_cancelled_proto_rawDescOnce.Do(func() {
		file_justeat_order_cancelled_proto_rawDescData = protoimpl.X.CompressGZIP(file_justeat_order_cancelled_proto_rawDescData)
	})
	return file_justeat_order_cancelled_proto_rawDescData
}

var file_justeat_order_cancelled_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_justeat_order_cancelled_proto_goTypes = []interface{}{
	(*OrderCancelled)(nil),                   // 0: justeat.OrderCancelled
	(*OrderCancelled_Transaction)(nil),       // 1: justeat.OrderCancelled.Transaction
	(*OrderCancelled_CancelationReason)(nil), // 2: justeat.OrderCancelled.CancelationReason
}
var file_justeat_order_cancelled_proto_depIdxs = []int32{
	1, // 0: justeat.OrderCancelled.customer_transaction:type_name -> justeat.OrderCancelled.Transaction
	2, // 1: justeat.OrderCancelled.reason:type_name -> justeat.OrderCancelled.CancelationReason
	1, // 2: justeat.OrderCancelled.restaurant_transaction:type_name -> justeat.OrderCancelled.Transaction
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_justeat_order_cancelled_proto_init() }
func file_justeat_order_cancelled_proto_init() {
	if File_justeat_order_cancelled_proto != nil {
		return
	}
	file_justeat_options_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_justeat_order_cancelled_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCancelled); i {
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
		file_justeat_order_cancelled_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCancelled_Transaction); i {
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
		file_justeat_order_cancelled_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCancelled_CancelationReason); i {
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
			RawDescriptor: file_justeat_order_cancelled_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_justeat_order_cancelled_proto_goTypes,
		DependencyIndexes: file_justeat_order_cancelled_proto_depIdxs,
		MessageInfos:      file_justeat_order_cancelled_proto_msgTypes,
	}.Build()
	File_justeat_order_cancelled_proto = out.File
	file_justeat_order_cancelled_proto_rawDesc = nil
	file_justeat_order_cancelled_proto_goTypes = nil
	file_justeat_order_cancelled_proto_depIdxs = nil
}
