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
	YourService_YourMethod_FullMethodName = "/apigrpc.YourService/YourMethod"
)

// YourServiceClient is the client API for YourService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type YourServiceClient interface {
	YourMethod(ctx context.Context, in *YourRequest, opts ...grpc.CallOption) (*YourResponse, error)
}

type yourServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewYourServiceClient(cc grpc.ClientConnInterface) YourServiceClient {
	return &yourServiceClient{cc}
}

func (c *yourServiceClient) YourMethod(ctx context.Context, in *YourRequest, opts ...grpc.CallOption) (*YourResponse, error) {
	out := new(YourResponse)
	err := c.cc.Invoke(ctx, YourService_YourMethod_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// YourServiceServer is the server API for YourService service.
// All implementations must embed UnimplementedYourServiceServer
// for forward compatibility
type YourServiceServer interface {
	YourMethod(context.Context, *YourRequest) (*YourResponse, error)
	mustEmbedUnimplementedYourServiceServer()
}

// UnimplementedYourServiceServer must be embedded to have forward compatible implementations.
type UnimplementedYourServiceServer struct {
}

func (UnimplementedYourServiceServer) YourMethod(context.Context, *YourRequest) (*YourResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method YourMethod not implemented")
}
func (UnimplementedYourServiceServer) mustEmbedUnimplementedYourServiceServer() {}

// UnsafeYourServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to YourServiceServer will
// result in compilation errors.
type UnsafeYourServiceServer interface {
	mustEmbedUnimplementedYourServiceServer()
}

func RegisterYourServiceServer(s grpc.ServiceRegistrar, srv YourServiceServer) {
	s.RegisterService(&YourService_ServiceDesc, srv)
}

func _YourService_YourMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(YourRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YourServiceServer).YourMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YourService_YourMethod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YourServiceServer).YourMethod(ctx, req.(*YourRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// YourService_ServiceDesc is the grpc.ServiceDesc for YourService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var YourService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apigrpc.YourService",
	HandlerType: (*YourServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "YourMethod",
			Handler:    _YourService_YourMethod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apigrpc.proto",
}
