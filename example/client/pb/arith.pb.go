// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: arith.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	rpc "net/rpc"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 算术运算请求结构
type ArithRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int32 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B int32 `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *ArithRequest) Reset() {
	*x = ArithRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arith_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArithRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArithRequest) ProtoMessage() {}

func (x *ArithRequest) ProtoReflect() protoreflect.Message {
	mi := &file_arith_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArithRequest.ProtoReflect.Descriptor instead.
func (*ArithRequest) Descriptor() ([]byte, []int) {
	return file_arith_proto_rawDescGZIP(), []int{0}
}

func (x *ArithRequest) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *ArithRequest) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

// 算术运算响应结构
type ArithResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pro int32 `protobuf:"varint,1,opt,name=pro,proto3" json:"pro,omitempty"` // 乘积
	Quo int32 `protobuf:"varint,2,opt,name=quo,proto3" json:"quo,omitempty"` // 商
	Rem int32 `protobuf:"varint,3,opt,name=rem,proto3" json:"rem,omitempty"` // 余数
}

func (x *ArithResponse) Reset() {
	*x = ArithResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arith_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArithResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArithResponse) ProtoMessage() {}

func (x *ArithResponse) ProtoReflect() protoreflect.Message {
	mi := &file_arith_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArithResponse.ProtoReflect.Descriptor instead.
func (*ArithResponse) Descriptor() ([]byte, []int) {
	return file_arith_proto_rawDescGZIP(), []int{1}
}

func (x *ArithResponse) GetPro() int32 {
	if x != nil {
		return x.Pro
	}
	return 0
}

func (x *ArithResponse) GetQuo() int32 {
	if x != nil {
		return x.Quo
	}
	return 0
}

func (x *ArithResponse) GetRem() int32 {
	if x != nil {
		return x.Rem
	}
	return 0
}

var File_arith_proto protoreflect.FileDescriptor

var file_arith_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x72, 0x69, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0x2a, 0x0a, 0x0c, 0x41, 0x72, 0x69, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x61, 0x12,
	0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x62, 0x22, 0x45, 0x0a,
	0x0d, 0x41, 0x72, 0x69, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x70, 0x72, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x72, 0x6f,
	0x12, 0x10, 0x0a, 0x03, 0x71, 0x75, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x71,
	0x75, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x72, 0x65, 0x6d, 0x32, 0x6e, 0x0a, 0x0c, 0x41, 0x72, 0x69, 0x74, 0x68, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79,
	0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x72, 0x69, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x72, 0x69, 0x74, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x12,
	0x10, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x72, 0x69, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x72, 0x69, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_arith_proto_rawDescOnce sync.Once
	file_arith_proto_rawDescData = file_arith_proto_rawDesc
)

func file_arith_proto_rawDescGZIP() []byte {
	file_arith_proto_rawDescOnce.Do(func() {
		file_arith_proto_rawDescData = protoimpl.X.CompressGZIP(file_arith_proto_rawDescData)
	})
	return file_arith_proto_rawDescData
}

var file_arith_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_arith_proto_goTypes = []interface{}{
	(*ArithRequest)(nil),  // 0: pb.ArithRequest
	(*ArithResponse)(nil), // 1: pb.ArithResponse
}
var file_arith_proto_depIdxs = []int32{
	0, // 0: pb.ArithService.multiply:input_type -> pb.ArithRequest
	0, // 1: pb.ArithService.divide:input_type -> pb.ArithRequest
	1, // 2: pb.ArithService.multiply:output_type -> pb.ArithResponse
	1, // 3: pb.ArithService.divide:output_type -> pb.ArithResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_arith_proto_init() }
func file_arith_proto_init() {
	if File_arith_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_arith_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArithRequest); i {
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
		file_arith_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArithResponse); i {
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
			RawDescriptor: file_arith_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_arith_proto_goTypes,
		DependencyIndexes: file_arith_proto_depIdxs,
		MessageInfos:      file_arith_proto_msgTypes,
	}.Build()
	File_arith_proto = out.File
	file_arith_proto_rawDesc = nil
	file_arith_proto_goTypes = nil
	file_arith_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ rpc.Server

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.

type ArithServiceService interface {
	Multiply(in *ArithRequest, out *ArithResponse) error
	Divide(in *ArithRequest, out *ArithResponse) error
}

type DefaultArithServiceService struct{}

func (d *DefaultArithServiceService) Multiply(in *ArithRequest, out *ArithResponse) error {
	panic("DefaultArithServiceService.Multiply is not implemented")
}

func (d *DefaultArithServiceService) Divide(in *ArithRequest, out *ArithResponse) error {
	panic("DefaultArithServiceService.Divide is not implemented")
}

// RegisterDefaultArithServiceServicepublish the given ArithService implementation on the server.
func RegisterDefaultArithServiceService(server *rpc.Server, i ArithServiceService) error {
	if err := server.RegisterName("ArithService", i); err != nil {
		return err
	}
	return nil
}

type ArithServiceClient struct {
	*rpc.Client
}

func (c *ArithServiceClient) Multiply(in *ArithRequest) (out *ArithResponse, err error) {
	if in == nil {
		in = new(ArithRequest)
	}

	out = new(ArithResponse)
	if err = c.Call("ArithService.Multiply", in, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ArithServiceClient) Divide(in *ArithRequest) (out *ArithResponse, err error) {
	if in == nil {
		in = new(ArithRequest)
	}

	out = new(ArithResponse)
	if err = c.Call("ArithService.Divide", in, out); err != nil {
		return nil, err
	}
	return out, nil
}