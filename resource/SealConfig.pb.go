//*
// @file:   SealConfig.proto
// @author: peng
// @brief:  这个文件是通过工具自动生成的，建议不要手动修改

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: SealConfig.proto

package resource

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

type SealConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//* 纹章ID
	SealID int32 `protobuf:"varint,1,opt,name=SealID,proto3" json:"SealID,omitempty"`
	//* 纹章名称
	SealName string `protobuf:"bytes,2,opt,name=SealName,proto3" json:"SealName,omitempty"`
	//* 属性
	SealAttr []*AttrConfig `protobuf:"bytes,3,rep,name=SealAttr,proto3" json:"SealAttr,omitempty"`
}

func (x *SealConfig) Reset() {
	*x = SealConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SealConfig_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SealConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SealConfig) ProtoMessage() {}

func (x *SealConfig) ProtoReflect() protoreflect.Message {
	mi := &file_SealConfig_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SealConfig.ProtoReflect.Descriptor instead.
func (*SealConfig) Descriptor() ([]byte, []int) {
	return file_SealConfig_proto_rawDescGZIP(), []int{0}
}

func (x *SealConfig) GetSealID() int32 {
	if x != nil {
		return x.SealID
	}
	return 0
}

func (x *SealConfig) GetSealName() string {
	if x != nil {
		return x.SealName
	}
	return ""
}

func (x *SealConfig) GetSealAttr() []*AttrConfig {
	if x != nil {
		return x.SealAttr
	}
	return nil
}

type AttrConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//* ID
	DataType int32 `protobuf:"varint,1,opt,name=DataType,proto3" json:"DataType,omitempty"`
	//* 类型
	AddType int32 `protobuf:"varint,2,opt,name=AddType,proto3" json:"AddType,omitempty"`
	//* 数值
	Value int32 `protobuf:"varint,3,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *AttrConfig) Reset() {
	*x = AttrConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SealConfig_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttrConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttrConfig) ProtoMessage() {}

func (x *AttrConfig) ProtoReflect() protoreflect.Message {
	mi := &file_SealConfig_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttrConfig.ProtoReflect.Descriptor instead.
func (*AttrConfig) Descriptor() ([]byte, []int) {
	return file_SealConfig_proto_rawDescGZIP(), []int{1}
}

func (x *AttrConfig) GetDataType() int32 {
	if x != nil {
		return x.DataType
	}
	return 0
}

func (x *AttrConfig) GetAddType() int32 {
	if x != nil {
		return x.AddType
	}
	return 0
}

func (x *AttrConfig) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type SealConfigList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*SealConfig `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *SealConfigList) Reset() {
	*x = SealConfigList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SealConfig_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SealConfigList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SealConfigList) ProtoMessage() {}

func (x *SealConfigList) ProtoReflect() protoreflect.Message {
	mi := &file_SealConfig_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SealConfigList.ProtoReflect.Descriptor instead.
func (*SealConfigList) Descriptor() ([]byte, []int) {
	return file_SealConfig_proto_rawDescGZIP(), []int{2}
}

func (x *SealConfigList) GetData() []*SealConfig {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_SealConfig_proto protoreflect.FileDescriptor

var file_SealConfig_proto_rawDesc = []byte{
	0x0a, 0x10, 0x53, 0x65, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x72, 0x0a, 0x0a,
	0x53, 0x65, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x65,
	0x61, 0x6c, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x65, 0x61, 0x6c,
	0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30,
	0x0a, 0x08, 0x53, 0x65, 0x61, 0x6c, 0x41, 0x74, 0x74, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x41, 0x74, 0x74, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x08, 0x53, 0x65, 0x61, 0x6c, 0x41, 0x74, 0x74, 0x72,
	0x22, 0x58, 0x0a, 0x0a, 0x41, 0x74, 0x74, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1a,
	0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x41, 0x64, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3a, 0x0a, 0x0e, 0x53, 0x65,
	0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x15, 0x5a, 0x13, 0x2e, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SealConfig_proto_rawDescOnce sync.Once
	file_SealConfig_proto_rawDescData = file_SealConfig_proto_rawDesc
)

func file_SealConfig_proto_rawDescGZIP() []byte {
	file_SealConfig_proto_rawDescOnce.Do(func() {
		file_SealConfig_proto_rawDescData = protoimpl.X.CompressGZIP(file_SealConfig_proto_rawDescData)
	})
	return file_SealConfig_proto_rawDescData
}

var file_SealConfig_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_SealConfig_proto_goTypes = []interface{}{
	(*SealConfig)(nil),     // 0: resource.SealConfig
	(*AttrConfig)(nil),     // 1: resource.AttrConfig
	(*SealConfigList)(nil), // 2: resource.SealConfigList
}
var file_SealConfig_proto_depIdxs = []int32{
	1, // 0: resource.SealConfig.SealAttr:type_name -> resource.AttrConfig
	0, // 1: resource.SealConfigList.data:type_name -> resource.SealConfig
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_SealConfig_proto_init() }
func file_SealConfig_proto_init() {
	if File_SealConfig_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SealConfig_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SealConfig); i {
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
		file_SealConfig_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttrConfig); i {
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
		file_SealConfig_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SealConfigList); i {
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
			RawDescriptor: file_SealConfig_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SealConfig_proto_goTypes,
		DependencyIndexes: file_SealConfig_proto_depIdxs,
		MessageInfos:      file_SealConfig_proto_msgTypes,
	}.Build()
	File_SealConfig_proto = out.File
	file_SealConfig_proto_rawDesc = nil
	file_SealConfig_proto_goTypes = nil
	file_SealConfig_proto_depIdxs = nil
}