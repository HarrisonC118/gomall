// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.29.3
// source: order_page.proto

package order

import (
	// _ "douyin-gomall/gomall/app/hertz_gateway/hertz_gen/api"
	// common "douyin-gomall/gomall/app/hertz_gateway/hertz_gen/hertz_gateway/common"
	_ "github.com/PiaoAdmin/gomall/app/hertz_gateway/hertz_gen/api"
	common "github.com/PiaoAdmin/gomall/app/hertz_gateway/hertz_gen/hertz_gateway/common"
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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_page_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_order_page_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_order_page_proto_rawDescGZIP(), []int{0}
}

var File_order_page_proto protoreflect.FileDescriptor

var file_order_page_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1a, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07,
	0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x5c, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x5f, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x1b, 0x2e, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x05,
	0xca, 0xc1, 0x18, 0x01, 0x2f, 0x42, 0x46, 0x5a, 0x44, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2d,
	0x67, 0x6f, 0x6d, 0x61, 0x6c, 0x6c, 0x2f, 0x67, 0x6f, 0x6d, 0x61, 0x6c, 0x6c, 0x2f, 0x61, 0x70,
	0x70, 0x2f, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f,
	0x68, 0x65, 0x72, 0x74, 0x7a, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x5f,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_page_proto_rawDescOnce sync.Once
	file_order_page_proto_rawDescData = file_order_page_proto_rawDesc
)

func file_order_page_proto_rawDescGZIP() []byte {
	file_order_page_proto_rawDescOnce.Do(func() {
		file_order_page_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_page_proto_rawDescData)
	})
	return file_order_page_proto_rawDescData
}

var file_order_page_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_order_page_proto_goTypes = []interface{}{
	(*Empty)(nil),        // 0: hertz_gateway.order.Empty
	(*common.Empty)(nil), // 1: hertz_gateway.common.Empty
}
var file_order_page_proto_depIdxs = []int32{
	1, // 0: hertz_gateway.order.OrderService.OrderList:input_type -> hertz_gateway.common.Empty
	1, // 1: hertz_gateway.order.OrderService.OrderList:output_type -> hertz_gateway.common.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_order_page_proto_init() }
func file_order_page_proto_init() {
	if File_order_page_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_page_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_order_page_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_page_proto_goTypes,
		DependencyIndexes: file_order_page_proto_depIdxs,
		MessageInfos:      file_order_page_proto_msgTypes,
	}.Build()
	File_order_page_proto = out.File
	file_order_page_proto_rawDesc = nil
	file_order_page_proto_goTypes = nil
	file_order_page_proto_depIdxs = nil
}
