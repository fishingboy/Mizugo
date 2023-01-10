// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: msgkey.proto

package messages

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

// 要求密鑰
type MsgKeyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgKeyReq) Reset() {
	*x = MsgKeyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msgkey_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgKeyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgKeyReq) ProtoMessage() {}

func (x *MsgKeyReq) ProtoReflect() protoreflect.Message {
	mi := &file_msgkey_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgKeyReq.ProtoReflect.Descriptor instead.
func (*MsgKeyReq) Descriptor() ([]byte, []int) {
	return file_msgkey_proto_rawDescGZIP(), []int{0}
}

// 回應密鑰
type MsgKeyRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"` // 密鑰
}

func (x *MsgKeyRes) Reset() {
	*x = MsgKeyRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msgkey_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgKeyRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgKeyRes) ProtoMessage() {}

func (x *MsgKeyRes) ProtoReflect() protoreflect.Message {
	mi := &file_msgkey_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgKeyRes.ProtoReflect.Descriptor instead.
func (*MsgKeyRes) Descriptor() ([]byte, []int) {
	return file_msgkey_proto_rawDescGZIP(), []int{1}
}

func (x *MsgKeyRes) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

var File_msgkey_proto protoreflect.FileDescriptor

var file_msgkey_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x73, 0x67, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0b,
	0x0a, 0x09, 0x4d, 0x73, 0x67, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x22, 0x1d, 0x0a, 0x09, 0x4d,
	0x73, 0x67, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x42, 0x14, 0x5a, 0x12, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x3b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msgkey_proto_rawDescOnce sync.Once
	file_msgkey_proto_rawDescData = file_msgkey_proto_rawDesc
)

func file_msgkey_proto_rawDescGZIP() []byte {
	file_msgkey_proto_rawDescOnce.Do(func() {
		file_msgkey_proto_rawDescData = protoimpl.X.CompressGZIP(file_msgkey_proto_rawDescData)
	})
	return file_msgkey_proto_rawDescData
}

var file_msgkey_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_msgkey_proto_goTypes = []interface{}{
	(*MsgKeyReq)(nil), // 0: MsgKeyReq
	(*MsgKeyRes)(nil), // 1: MsgKeyRes
}
var file_msgkey_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_msgkey_proto_init() }
func file_msgkey_proto_init() {
	if File_msgkey_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msgkey_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgKeyReq); i {
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
		file_msgkey_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgKeyRes); i {
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
			RawDescriptor: file_msgkey_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msgkey_proto_goTypes,
		DependencyIndexes: file_msgkey_proto_depIdxs,
		MessageInfos:      file_msgkey_proto_msgTypes,
	}.Build()
	File_msgkey_proto = out.File
	file_msgkey_proto_rawDesc = nil
	file_msgkey_proto_goTypes = nil
	file_msgkey_proto_depIdxs = nil
}
