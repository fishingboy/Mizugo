// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: msgverify.proto

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
		mi := &file_msgverify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgKeyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgKeyReq) ProtoMessage() {}

func (x *MsgKeyReq) ProtoReflect() protoreflect.Message {
	mi := &file_msgverify_proto_msgTypes[0]
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
	return file_msgverify_proto_rawDescGZIP(), []int{0}
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
		mi := &file_msgverify_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgKeyRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgKeyRes) ProtoMessage() {}

func (x *MsgKeyRes) ProtoReflect() protoreflect.Message {
	mi := &file_msgverify_proto_msgTypes[1]
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
	return file_msgverify_proto_rawDescGZIP(), []int{1}
}

func (x *MsgKeyRes) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// 要求Ping
type MsgPingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time int64 `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"` // 傳送時間
}

func (x *MsgPingReq) Reset() {
	*x = MsgPingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msgverify_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgPingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgPingReq) ProtoMessage() {}

func (x *MsgPingReq) ProtoReflect() protoreflect.Message {
	mi := &file_msgverify_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgPingReq.ProtoReflect.Descriptor instead.
func (*MsgPingReq) Descriptor() ([]byte, []int) {
	return file_msgverify_proto_rawDescGZIP(), []int{2}
}

func (x *MsgPingReq) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

// 回應Ping
type MsgPingRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From  *MsgPingReq `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`    // 來源訊息
	Count int64       `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"` // 封包計數
}

func (x *MsgPingRes) Reset() {
	*x = MsgPingRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msgverify_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgPingRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgPingRes) ProtoMessage() {}

func (x *MsgPingRes) ProtoReflect() protoreflect.Message {
	mi := &file_msgverify_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgPingRes.ProtoReflect.Descriptor instead.
func (*MsgPingRes) Descriptor() ([]byte, []int) {
	return file_msgverify_proto_rawDescGZIP(), []int{3}
}

func (x *MsgPingRes) GetFrom() *MsgPingReq {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *MsgPingRes) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_msgverify_proto protoreflect.FileDescriptor

var file_msgverify_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6d, 0x73, 0x67, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x0b, 0x0a, 0x09, 0x4d, 0x73, 0x67, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x22, 0x1d,
	0x0a, 0x09, 0x4d, 0x73, 0x67, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x20, 0x0a,
	0x0a, 0x4d, 0x73, 0x67, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22,
	0x43, 0x0a, 0x0a, 0x4d, 0x73, 0x67, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x12, 0x1f, 0x0a,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4d, 0x73,
	0x67, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x42, 0x14, 0x5a, 0x12, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x3b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_msgverify_proto_rawDescOnce sync.Once
	file_msgverify_proto_rawDescData = file_msgverify_proto_rawDesc
)

func file_msgverify_proto_rawDescGZIP() []byte {
	file_msgverify_proto_rawDescOnce.Do(func() {
		file_msgverify_proto_rawDescData = protoimpl.X.CompressGZIP(file_msgverify_proto_rawDescData)
	})
	return file_msgverify_proto_rawDescData
}

var file_msgverify_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_msgverify_proto_goTypes = []interface{}{
	(*MsgKeyReq)(nil),  // 0: MsgKeyReq
	(*MsgKeyRes)(nil),  // 1: MsgKeyRes
	(*MsgPingReq)(nil), // 2: MsgPingReq
	(*MsgPingRes)(nil), // 3: MsgPingRes
}
var file_msgverify_proto_depIdxs = []int32{
	2, // 0: MsgPingRes.from:type_name -> MsgPingReq
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_msgverify_proto_init() }
func file_msgverify_proto_init() {
	if File_msgverify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msgverify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_msgverify_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_msgverify_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgPingReq); i {
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
		file_msgverify_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgPingRes); i {
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
			RawDescriptor: file_msgverify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msgverify_proto_goTypes,
		DependencyIndexes: file_msgverify_proto_depIdxs,
		MessageInfos:      file_msgverify_proto_msgTypes,
	}.Build()
	File_msgverify_proto = out.File
	file_msgverify_proto_rawDesc = nil
	file_msgverify_proto_goTypes = nil
	file_msgverify_proto_depIdxs = nil
}
