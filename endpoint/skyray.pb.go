// Code generated by protoc-gen-go. DO NOT EDIT.
// source: skyray.proto

/*
Package endpoint is a generated protocol buffer package.

It is generated from these files:
	skyray.proto

It has these top-level messages:
	Command
	Response
*/
package endpoint

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

type Command struct {
	Tokens []string `protobuf:"bytes,1,rep,name=tokens" json:"tokens,omitempty"`
}

func (m *Command) Reset()                    { *m = Command{} }
func (m *Command) String() string            { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()               {}
func (*Command) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Command) GetTokens() []string {
	if m != nil {
		return m.Tokens
	}
	return nil
}

type Response struct {
	Output string `protobuf:"bytes,1,opt,name=output" json:"output,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func init() {
	proto.RegisterType((*Command)(nil), "endpoint.Command")
	proto.RegisterType((*Response)(nil), "endpoint.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SkyrayService service

type SkyrayServiceClient interface {
	Connect(ctx context.Context, opts ...grpc.CallOption) (SkyrayService_ConnectClient, error)
}

type skyrayServiceClient struct {
	cc *grpc.ClientConn
}

func NewSkyrayServiceClient(cc *grpc.ClientConn) SkyrayServiceClient {
	return &skyrayServiceClient{cc}
}

func (c *skyrayServiceClient) Connect(ctx context.Context, opts ...grpc.CallOption) (SkyrayService_ConnectClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SkyrayService_serviceDesc.Streams[0], c.cc, "/endpoint.SkyrayService/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &skyrayServiceConnectClient{stream}
	return x, nil
}

type SkyrayService_ConnectClient interface {
	Send(*Command) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type skyrayServiceConnectClient struct {
	grpc.ClientStream
}

func (x *skyrayServiceConnectClient) Send(m *Command) error {
	return x.ClientStream.SendMsg(m)
}

func (x *skyrayServiceConnectClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for SkyrayService service

type SkyrayServiceServer interface {
	Connect(SkyrayService_ConnectServer) error
}

func RegisterSkyrayServiceServer(s *grpc.Server, srv SkyrayServiceServer) {
	s.RegisterService(&_SkyrayService_serviceDesc, srv)
}

func _SkyrayService_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SkyrayServiceServer).Connect(&skyrayServiceConnectServer{stream})
}

type SkyrayService_ConnectServer interface {
	Send(*Response) error
	Recv() (*Command, error)
	grpc.ServerStream
}

type skyrayServiceConnectServer struct {
	grpc.ServerStream
}

func (x *skyrayServiceConnectServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *skyrayServiceConnectServer) Recv() (*Command, error) {
	m := new(Command)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SkyrayService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "endpoint.SkyrayService",
	HandlerType: (*SkyrayServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _SkyrayService_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "skyray.proto",
}

func init() { proto.RegisterFile("skyray.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0xce, 0xae, 0x2c,
	0x4a, 0xac, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x48, 0xcd, 0x4b, 0x29, 0xc8, 0xcf,
	0xcc, 0x2b, 0x51, 0x52, 0xe4, 0x62, 0x77, 0xce, 0xcf, 0xcd, 0x4d, 0xcc, 0x4b, 0x11, 0x12, 0xe3,
	0x62, 0x2b, 0xc9, 0xcf, 0x4e, 0xcd, 0x2b, 0x96, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x0c, 0x82, 0xf2,
	0x94, 0x94, 0xb8, 0x38, 0x82, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x41, 0x6a, 0xf2, 0x4b,
	0x4b, 0x0a, 0x4a, 0x4b, 0x24, 0x18, 0x15, 0x18, 0x41, 0x6a, 0x20, 0x3c, 0x23, 0x77, 0x2e, 0xde,
	0x60, 0xb0, 0x05, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0x66, 0x20, 0x73, 0xf3, 0xf2,
	0x52, 0x93, 0x4b, 0x84, 0x04, 0xf5, 0x60, 0xb6, 0xe9, 0x41, 0xad, 0x92, 0x12, 0x42, 0x08, 0xc1,
	0x8c, 0x56, 0x62, 0xd0, 0x60, 0x34, 0x60, 0x4c, 0x62, 0x03, 0x3b, 0xd0, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0x25, 0xf9, 0x76, 0x68, 0xb0, 0x00, 0x00, 0x00,
}