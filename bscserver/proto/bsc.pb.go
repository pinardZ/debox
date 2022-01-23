// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: bsc.proto

package proto

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

type ReqAccountBalance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *ReqAccountBalance) Reset() {
	*x = ReqAccountBalance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bsc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqAccountBalance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqAccountBalance) ProtoMessage() {}

func (x *ReqAccountBalance) ProtoReflect() protoreflect.Message {
	mi := &file_bsc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqAccountBalance.ProtoReflect.Descriptor instead.
func (*ReqAccountBalance) Descriptor() ([]byte, []int) {
	return file_bsc_proto_rawDescGZIP(), []int{0}
}

func (x *ReqAccountBalance) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type RespAccountBalance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Result  string `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *RespAccountBalance) Reset() {
	*x = RespAccountBalance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bsc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespAccountBalance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespAccountBalance) ProtoMessage() {}

func (x *RespAccountBalance) ProtoReflect() protoreflect.Message {
	mi := &file_bsc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespAccountBalance.ProtoReflect.Descriptor instead.
func (*RespAccountBalance) Descriptor() ([]byte, []int) {
	return file_bsc_proto_rawDescGZIP(), []int{1}
}

func (x *RespAccountBalance) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RespAccountBalance) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RespAccountBalance) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_bsc_proto protoreflect.FileDescriptor

var file_bsc_proto_rawDesc = []byte{
	0x0a, 0x09, 0x62, 0x73, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x22, 0x5a, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x55, 0x0a,
	0x0a, 0x42, 0x53, 0x43, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bsc_proto_rawDescOnce sync.Once
	file_bsc_proto_rawDescData = file_bsc_proto_rawDesc
)

func file_bsc_proto_rawDescGZIP() []byte {
	file_bsc_proto_rawDescOnce.Do(func() {
		file_bsc_proto_rawDescData = protoimpl.X.CompressGZIP(file_bsc_proto_rawDescData)
	})
	return file_bsc_proto_rawDescData
}

var file_bsc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bsc_proto_goTypes = []interface{}{
	(*ReqAccountBalance)(nil),  // 0: proto.ReqAccountBalance
	(*RespAccountBalance)(nil), // 1: proto.RespAccountBalance
}
var file_bsc_proto_depIdxs = []int32{
	0, // 0: proto.BSCService.AccountBalance:input_type -> proto.ReqAccountBalance
	1, // 1: proto.BSCService.AccountBalance:output_type -> proto.RespAccountBalance
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bsc_proto_init() }
func file_bsc_proto_init() {
	if File_bsc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bsc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqAccountBalance); i {
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
		file_bsc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespAccountBalance); i {
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
			RawDescriptor: file_bsc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bsc_proto_goTypes,
		DependencyIndexes: file_bsc_proto_depIdxs,
		MessageInfos:      file_bsc_proto_msgTypes,
	}.Build()
	File_bsc_proto = out.File
	file_bsc_proto_rawDesc = nil
	file_bsc_proto_goTypes = nil
	file_bsc_proto_depIdxs = nil
}
