// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: proto/ping-pong.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PingPongService_Ping_FullMethodName = "/proto.PingPongService/Ping"
)

// PingPongServiceClient is the client API for PingPongService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PingPongServiceClient interface {
	Ping(ctx context.Context, in *PingPongRequest, opts ...grpc.CallOption) (*PingPongReply, error)
}

type pingPongServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPingPongServiceClient(cc grpc.ClientConnInterface) PingPongServiceClient {
	return &pingPongServiceClient{cc}
}

func (c *pingPongServiceClient) Ping(ctx context.Context, in *PingPongRequest, opts ...grpc.CallOption) (*PingPongReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PingPongReply)
	err := c.cc.Invoke(ctx, PingPongService_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PingPongServiceServer is the server API for PingPongService service.
// All implementations must embed UnimplementedPingPongServiceServer
// for forward compatibility.
type PingPongServiceServer interface {
	Ping(context.Context, *PingPongRequest) (*PingPongReply, error)
	mustEmbedUnimplementedPingPongServiceServer()
}

// UnimplementedPingPongServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPingPongServiceServer struct{}

func (UnimplementedPingPongServiceServer) Ping(context.Context, *PingPongRequest) (*PingPongReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedPingPongServiceServer) mustEmbedUnimplementedPingPongServiceServer() {}
func (UnimplementedPingPongServiceServer) testEmbeddedByValue()                         {}

// UnsafePingPongServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PingPongServiceServer will
// result in compilation errors.
type UnsafePingPongServiceServer interface {
	mustEmbedUnimplementedPingPongServiceServer()
}

func RegisterPingPongServiceServer(s grpc.ServiceRegistrar, srv PingPongServiceServer) {
	// If the following call pancis, it indicates UnimplementedPingPongServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PingPongService_ServiceDesc, srv)
}

func _PingPongService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingPongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingPongServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PingPongService_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingPongServiceServer).Ping(ctx, req.(*PingPongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PingPongService_ServiceDesc is the grpc.ServiceDesc for PingPongService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PingPongService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PingPongService",
	HandlerType: (*PingPongServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _PingPongService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/ping-pong.proto",
}
