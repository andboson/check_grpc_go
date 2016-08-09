// Code generated by protoc-gen-go.
// source: helloworld.proto
// DO NOT EDIT!

/*
Package helloworld is a generated protocol buffer package.

It is generated from these files:
	helloworld.proto

It has these top-level messages:
	HelloRequest1
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
type HelloRequest1 struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	WaitMinutes int32  `protobuf:"varint,2,opt,name=wait_minutes,json=waitMinutes" json:"wait_minutes,omitempty"`
}

func (m *HelloRequest1) Reset()                    { *m = HelloRequest1{} }
func (m *HelloRequest1) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest1) ProtoMessage()               {}
func (*HelloRequest1) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// The response message containing the greetings
type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*HelloRequest1)(nil), "helloworld.HelloRequest1")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Greeter service

type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest1, opts ...grpc.CallOption) (*HelloReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest1, opts ...grpc.CallOption) (*HelloReply, error) {
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
	SayHello(context.Context, *HelloRequest1) (*HelloReply, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest1)
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
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest1))
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
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x28, 0xb9, 0x71, 0xf1, 0x7a, 0x80, 0x78, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x86, 0x42,
	0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6,
	0x90, 0x22, 0x17, 0x4f, 0x79, 0x62, 0x66, 0x49, 0x7c, 0x6e, 0x66, 0x5e, 0x69, 0x49, 0x6a, 0xb1,
	0x04, 0x93, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x37, 0x48, 0xcc, 0x17, 0x22, 0xa4, 0xa4, 0xc6, 0xc5,
	0x05, 0x35, 0xa7, 0x20, 0xa7, 0x52, 0x48, 0x82, 0x8b, 0x3d, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x1d,
	0x66, 0x0e, 0x8c, 0x6b, 0xe4, 0xc5, 0xc5, 0xee, 0x5e, 0x94, 0x9a, 0x5a, 0x92, 0x5a, 0x24, 0x64,
	0xcf, 0xc5, 0x11, 0x9c, 0x58, 0x09, 0xd6, 0x25, 0x24, 0xa9, 0x87, 0xe4, 0x4a, 0x14, 0x07, 0x49,
	0x89, 0x61, 0x91, 0x2a, 0xc8, 0xa9, 0x54, 0x62, 0x70, 0x32, 0xe0, 0x92, 0xce, 0xcc, 0xd7, 0x4b,
	0x2f, 0x2a, 0x48, 0xd6, 0x4b, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0x2d, 0x46, 0x52, 0xeb, 0xc4,
	0x0f, 0x56, 0x1c, 0x0e, 0x62, 0x07, 0x80, 0xfc, 0x1d, 0xc0, 0x98, 0xc4, 0x06, 0x0e, 0x00, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x1b, 0x5e, 0x85, 0x14, 0x01, 0x00, 0x00,
}
