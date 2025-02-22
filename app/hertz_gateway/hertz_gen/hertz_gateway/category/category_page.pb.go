// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.29.3
// source: category_page.proto

package category

import (
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
		mi := &file_category_page_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_category_page_proto_msgTypes[0]
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
	return file_category_page_proto_rawDescGZIP(), []int{0}
}

type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" form:"id" query:"id"`
	CategoryName string `protobuf:"bytes,2,opt,name=categoryName,proto3" json:"categoryName,omitempty" form:"categoryName" query:"categoryName"`
	Status       int32  `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty" form:"status" query:"status"`
	ParentId     int64  `protobuf:"varint,5,opt,name=parentId,proto3" json:"parentId,omitempty" form:"parentId" query:"parentId"`
}

func (x *Category) Reset() {
	*x = Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_page_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_category_page_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_category_page_proto_rawDescGZIP(), []int{1}
}

func (x *Category) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Category) GetCategoryName() string {
	if x != nil {
		return x.CategoryName
	}
	return ""
}

func (x *Category) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Category) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

type AddCategoryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category *Category `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty" form:"category" query:"category"`
}

func (x *AddCategoryReq) Reset() {
	*x = AddCategoryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_page_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCategoryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCategoryReq) ProtoMessage() {}

func (x *AddCategoryReq) ProtoReflect() protoreflect.Message {
	mi := &file_category_page_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCategoryReq.ProtoReflect.Descriptor instead.
func (*AddCategoryReq) Descriptor() ([]byte, []int) {
	return file_category_page_proto_rawDescGZIP(), []int{2}
}

func (x *AddCategoryReq) GetCategory() *Category {
	if x != nil {
		return x.Category
	}
	return nil
}

type CategoryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category string `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty" path:"category"`
}

func (x *CategoryReq) Reset() {
	*x = CategoryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_page_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryReq) ProtoMessage() {}

func (x *CategoryReq) ProtoReflect() protoreflect.Message {
	mi := &file_category_page_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryReq.ProtoReflect.Descriptor instead.
func (*CategoryReq) Descriptor() ([]byte, []int) {
	return file_category_page_proto_rawDescGZIP(), []int{3}
}

func (x *CategoryReq) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

var File_category_page_proto protoreflect.FileDescriptor

var file_category_page_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x5f, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1a, 0x09, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x5f,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x72, 0x0a,
	0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x22, 0x4e, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x12, 0x3c, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x5f, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x22, 0x37, 0x0a, 0x0b, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x12, 0x28, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0c, 0xd2, 0xbb, 0x18, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x32, 0xe2, 0x01, 0x0a, 0x0f, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x65,
	0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x23, 0x2e, 0x68, 0x65, 0x72,
	0x74, 0x7a, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x1a,
	0x1b, 0x2e, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x17, 0xca, 0xc1,
	0x18, 0x13, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2f, 0x3a, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x68, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x12, 0x26, 0x2e, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x5f, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x41, 0x64,
	0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x68,
	0x65, 0x72, 0x74, 0x7a, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x14, 0xd2, 0xc1, 0x18, 0x10, 0x2f,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x50, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x69,
	0x61, 0x6f, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x67, 0x6f, 0x6d, 0x61, 0x6c, 0x6c, 0x2f, 0x61,
	0x70, 0x70, 0x2f, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2f, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x68, 0x65, 0x72, 0x74, 0x7a,
	0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_category_page_proto_rawDescOnce sync.Once
	file_category_page_proto_rawDescData = file_category_page_proto_rawDesc
)

func file_category_page_proto_rawDescGZIP() []byte {
	file_category_page_proto_rawDescOnce.Do(func() {
		file_category_page_proto_rawDescData = protoimpl.X.CompressGZIP(file_category_page_proto_rawDescData)
	})
	return file_category_page_proto_rawDescData
}

var file_category_page_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_category_page_proto_goTypes = []interface{}{
	(*Empty)(nil),          // 0: hertz_gateway.category.Empty
	(*Category)(nil),       // 1: hertz_gateway.category.Category
	(*AddCategoryReq)(nil), // 2: hertz_gateway.category.AddCategoryReq
	(*CategoryReq)(nil),    // 3: hertz_gateway.category.CategoryReq
	(*common.Empty)(nil),   // 4: hertz_gateway.common.Empty
}
var file_category_page_proto_depIdxs = []int32{
	1, // 0: hertz_gateway.category.AddCategoryReq.category:type_name -> hertz_gateway.category.Category
	3, // 1: hertz_gateway.category.CategoryService.Category:input_type -> hertz_gateway.category.CategoryReq
	2, // 2: hertz_gateway.category.CategoryService.AddCategory:input_type -> hertz_gateway.category.AddCategoryReq
	4, // 3: hertz_gateway.category.CategoryService.Category:output_type -> hertz_gateway.common.Empty
	4, // 4: hertz_gateway.category.CategoryService.AddCategory:output_type -> hertz_gateway.common.Empty
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_category_page_proto_init() }
func file_category_page_proto_init() {
	if File_category_page_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_category_page_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_category_page_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Category); i {
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
		file_category_page_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCategoryReq); i {
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
		file_category_page_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryReq); i {
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
			RawDescriptor: file_category_page_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_category_page_proto_goTypes,
		DependencyIndexes: file_category_page_proto_depIdxs,
		MessageInfos:      file_category_page_proto_msgTypes,
	}.Build()
	File_category_page_proto = out.File
	file_category_page_proto_rawDesc = nil
	file_category_page_proto_goTypes = nil
	file_category_page_proto_depIdxs = nil
}
