// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld/helloworld.proto

/*
Package helloworld is a generated protocol buffer package.

It is generated from these files:
	helloworld/helloworld.proto

It has these top-level messages:
	HelloRequest
	HelloReply
*/
package helloworld

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request message containing the user's name.
type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Code    int64  `protobuf:"varint,10010,opt,name=code" json:"code,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *HelloReply) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Greeter service

type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := grpc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld/helloworld.proto",
}

func init() { proto.RegisterFile("helloworld/helloworld.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x47, 0x30, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85,
	0xb8, 0x10, 0x22, 0x4a, 0x4a, 0x5c, 0x3c, 0x1e, 0x20, 0x5e, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71,
	0x89, 0x90, 0x10, 0x17, 0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10,
	0x98, 0xad, 0x64, 0xcd, 0xc5, 0x05, 0x55, 0x53, 0x90, 0x53, 0x29, 0x24, 0xc1, 0xc5, 0x9e, 0x9b,
	0x5a, 0x5c, 0x9c, 0x98, 0x0e, 0x53, 0x04, 0xe3, 0x0a, 0x09, 0x73, 0xb1, 0x24, 0xe7, 0xa7, 0xa4,
	0x4a, 0xcc, 0xf2, 0x53, 0x60, 0xd4, 0x60, 0x0e, 0x02, 0x73, 0x8c, 0x3c, 0xb9, 0xd8, 0xdd, 0x8b,
	0x52, 0x53, 0x4b, 0x52, 0x8b, 0x84, 0xec, 0xb8, 0x38, 0x82, 0x13, 0x2b, 0xc1, 0x46, 0x09, 0x49,
	0xe8, 0x21, 0x39, 0x0b, 0xd9, 0x05, 0x52, 0x62, 0x58, 0x64, 0x0a, 0x72, 0x2a, 0x95, 0x18, 0x9c,
	0x0c, 0xb8, 0xa4, 0x33, 0xf3, 0xf5, 0xd2, 0x8b, 0x0a, 0x92, 0xf5, 0x52, 0x2b, 0x12, 0x73, 0x0b,
	0x72, 0x52, 0x8b, 0x91, 0xd4, 0x3a, 0xf1, 0x83, 0x15, 0x87, 0x83, 0xd8, 0x01, 0x20, 0x7f, 0x06,
	0x30, 0x26, 0xb1, 0x81, 0x3d, 0x6c, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x57, 0x8e, 0x82,
	0x0f, 0x01, 0x00, 0x00,
}