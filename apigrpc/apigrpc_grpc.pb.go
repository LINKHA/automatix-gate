// apigrpc.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: apigrpc.proto

package apigrpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GateServer_SayHello_FullMethodName = "/apigrpc.GateServer/SayHello"
)

// GateServerClient is the client API for GateServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GateServerClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type gateServerClient struct {
	cc grpc.ClientConnInterface
}

func NewGateServerClient(cc grpc.ClientConnInterface) GateServerClient {
	return &gateServerClient{cc}
}

func (c *gateServerClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, GateServer_SayHello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GateServerServer is the server API for GateServer service.
// All implementations must embed UnimplementedGateServerServer
// for forward compatibility
type GateServerServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedGateServerServer()
}

// UnimplementedGateServerServer must be embedded to have forward compatible implementations.
type UnimplementedGateServerServer struct {
}

func (UnimplementedGateServerServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGateServerServer) mustEmbedUnimplementedGateServerServer() {}

// UnsafeGateServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GateServerServer will
// result in compilation errors.
type UnsafeGateServerServer interface {
	mustEmbedUnimplementedGateServerServer()
}

func RegisterGateServerServer(s grpc.ServiceRegistrar, srv GateServerServer) {
	s.RegisterService(&GateServer_ServiceDesc, srv)
}

func _GateServer_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GateServerServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GateServer_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GateServerServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GateServer_ServiceDesc is the grpc.ServiceDesc for GateServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GateServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apigrpc.GateServer",
	HandlerType: (*GateServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _GateServer_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apigrpc.proto",
}
