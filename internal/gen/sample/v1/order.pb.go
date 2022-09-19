// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: sample/v1/order.proto

package samplev1

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

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId   string      `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	LineItems []*LineItem `protobuf:"bytes,2,rep,name=line_items,json=lineItems,proto3" json:"line_items,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sample_v1_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_sample_v1_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_sample_v1_order_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Order) GetLineItems() []*LineItem {
	if x != nil {
		return x.LineItems
	}
	return nil
}

type LineItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LineItemId string   `protobuf:"bytes,1,opt,name=line_item_id,json=lineItemId,proto3" json:"line_item_id,omitempty"`
	BasePrice  int64    `protobuf:"varint,2,opt,name=base_price,json=basePrice,proto3" json:"base_price,omitempty"`
	Markup     float64  `protobuf:"fixed64,3,opt,name=markup,proto3" json:"markup,omitempty"`
	IsPromo    bool     `protobuf:"varint,4,opt,name=is_promo,json=isPromo,proto3" json:"is_promo,omitempty"`
	Quantity   int64    `protobuf:"varint,5,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Address    *Address `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *LineItem) Reset() {
	*x = LineItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sample_v1_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LineItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineItem) ProtoMessage() {}

func (x *LineItem) ProtoReflect() protoreflect.Message {
	mi := &file_sample_v1_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LineItem.ProtoReflect.Descriptor instead.
func (*LineItem) Descriptor() ([]byte, []int) {
	return file_sample_v1_order_proto_rawDescGZIP(), []int{1}
}

func (x *LineItem) GetLineItemId() string {
	if x != nil {
		return x.LineItemId
	}
	return ""
}

func (x *LineItem) GetBasePrice() int64 {
	if x != nil {
		return x.BasePrice
	}
	return 0
}

func (x *LineItem) GetMarkup() float64 {
	if x != nil {
		return x.Markup
	}
	return 0
}

func (x *LineItem) GetIsPromo() bool {
	if x != nil {
		return x.IsPromo
	}
	return false
}

func (x *LineItem) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *LineItem) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DetailAddress string `protobuf:"bytes,1,opt,name=detail_address,json=detailAddress,proto3" json:"detail_address,omitempty"`
	City          string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sample_v1_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_sample_v1_order_proto_msgTypes[2]
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
	return file_sample_v1_order_proto_rawDescGZIP(), []int{2}
}

func (x *Address) GetDetailAddress() string {
	if x != nil {
		return x.DetailAddress
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

var File_sample_v1_order_proto protoreflect.FileDescriptor

var file_sample_v1_order_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x22, 0x56, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x0a, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x09, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xc8, 0x01, 0x0a, 0x08, 0x4c,
	0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x20, 0x0a, 0x0c, 0x6c, 0x69, 0x6e, 0x65, 0x5f,
	0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c,
	0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62,
	0x61, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61, 0x72, 0x6b,
	0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x75, 0x70,
	0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x2c, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x44, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x42, 0xa1, 0x01, 0x0a, 0x0d,
	0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2d, 0x64,
	0x65, 0x76, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x61, 0x74, 0x68, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2f, 0x76, 0x31, 0x3b, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53,
	0x58, 0x58, 0xaa, 0x02, 0x09, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x09, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x53, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x0a, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sample_v1_order_proto_rawDescOnce sync.Once
	file_sample_v1_order_proto_rawDescData = file_sample_v1_order_proto_rawDesc
)

func file_sample_v1_order_proto_rawDescGZIP() []byte {
	file_sample_v1_order_proto_rawDescOnce.Do(func() {
		file_sample_v1_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_sample_v1_order_proto_rawDescData)
	})
	return file_sample_v1_order_proto_rawDescData
}

var file_sample_v1_order_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sample_v1_order_proto_goTypes = []interface{}{
	(*Order)(nil),    // 0: sample.v1.Order
	(*LineItem)(nil), // 1: sample.v1.LineItem
	(*Address)(nil),  // 2: sample.v1.Address
}
var file_sample_v1_order_proto_depIdxs = []int32{
	1, // 0: sample.v1.Order.line_items:type_name -> sample.v1.LineItem
	2, // 1: sample.v1.LineItem.address:type_name -> sample.v1.Address
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_sample_v1_order_proto_init() }
func file_sample_v1_order_proto_init() {
	if File_sample_v1_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sample_v1_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_sample_v1_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LineItem); i {
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
		file_sample_v1_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_sample_v1_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sample_v1_order_proto_goTypes,
		DependencyIndexes: file_sample_v1_order_proto_depIdxs,
		MessageInfos:      file_sample_v1_order_proto_msgTypes,
	}.Build()
	File_sample_v1_order_proto = out.File
	file_sample_v1_order_proto_rawDesc = nil
	file_sample_v1_order_proto_goTypes = nil
	file_sample_v1_order_proto_depIdxs = nil
}
