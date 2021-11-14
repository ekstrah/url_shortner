// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package go_urlBox_grpc

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

// GenURLManagementClient is the client API for GenURLManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GenURLManagementClient interface {
	GenNewURL(ctx context.Context, in *ExURLReq, opts ...grpc.CallOption) (*ExURLRes, error)
	ReDirURL(ctx context.Context, in *ReDirReq, opts ...grpc.CallOption) (*ReDirRes, error)
}

type genURLManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewGenURLManagementClient(cc grpc.ClientConnInterface) GenURLManagementClient {
	return &genURLManagementClient{cc}
}

func (c *genURLManagementClient) GenNewURL(ctx context.Context, in *ExURLReq, opts ...grpc.CallOption) (*ExURLRes, error) {
	out := new(ExURLRes)
	err := c.cc.Invoke(ctx, "/urlBox.GenURLManagement/GenNewURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genURLManagementClient) ReDirURL(ctx context.Context, in *ReDirReq, opts ...grpc.CallOption) (*ReDirRes, error) {
	out := new(ReDirRes)
	err := c.cc.Invoke(ctx, "/urlBox.GenURLManagement/ReDirURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GenURLManagementServer is the server API for GenURLManagement service.
// All implementations must embed UnimplementedGenURLManagementServer
// for forward compatibility
type GenURLManagementServer interface {
	GenNewURL(context.Context, *ExURLReq) (*ExURLRes, error)
	ReDirURL(context.Context, *ReDirReq) (*ReDirRes, error)
	mustEmbedUnimplementedGenURLManagementServer()
}

// UnimplementedGenURLManagementServer must be embedded to have forward compatible implementations.
type UnimplementedGenURLManagementServer struct {
}

func (UnimplementedGenURLManagementServer) GenNewURL(context.Context, *ExURLReq) (*ExURLRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenNewURL not implemented")
}
func (UnimplementedGenURLManagementServer) ReDirURL(context.Context, *ReDirReq) (*ReDirRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReDirURL not implemented")
}
func (UnimplementedGenURLManagementServer) mustEmbedUnimplementedGenURLManagementServer() {}

// UnsafeGenURLManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GenURLManagementServer will
// result in compilation errors.
type UnsafeGenURLManagementServer interface {
	mustEmbedUnimplementedGenURLManagementServer()
}

func RegisterGenURLManagementServer(s grpc.ServiceRegistrar, srv GenURLManagementServer) {
	s.RegisterService(&GenURLManagement_ServiceDesc, srv)
}

func _GenURLManagement_GenNewURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExURLReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenURLManagementServer).GenNewURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/urlBox.GenURLManagement/GenNewURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenURLManagementServer).GenNewURL(ctx, req.(*ExURLReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GenURLManagement_ReDirURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReDirReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenURLManagementServer).ReDirURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/urlBox.GenURLManagement/ReDirURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenURLManagementServer).ReDirURL(ctx, req.(*ReDirReq))
	}
	return interceptor(ctx, in, info, handler)
}

// GenURLManagement_ServiceDesc is the grpc.ServiceDesc for GenURLManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GenURLManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "urlBox.GenURLManagement",
	HandlerType: (*GenURLManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenNewURL",
			Handler:    _GenURLManagement_GenNewURL_Handler,
		},
		{
			MethodName: "ReDirURL",
			Handler:    _GenURLManagement_ReDirURL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "urlBox.proto",
}
