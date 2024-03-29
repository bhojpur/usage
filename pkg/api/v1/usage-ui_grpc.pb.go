// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// UsageUIClient is the client API for UsageUI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsageUIClient interface {
	// ListEngineSpecs returns a list of Usage Engine(s) that can be started through the UI.
	ListEngineSpecs(ctx context.Context, in *ListEngineSpecsRequest, opts ...grpc.CallOption) (UsageUI_ListEngineSpecsClient, error)
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error)
}

type usageUIClient struct {
	cc grpc.ClientConnInterface
}

func NewUsageUIClient(cc grpc.ClientConnInterface) UsageUIClient {
	return &usageUIClient{cc}
}

func (c *usageUIClient) ListEngineSpecs(ctx context.Context, in *ListEngineSpecsRequest, opts ...grpc.CallOption) (UsageUI_ListEngineSpecsClient, error) {
	stream, err := c.cc.NewStream(ctx, &UsageUI_ServiceDesc.Streams[0], "/v1.UsageUI/ListEngineSpecs", opts...)
	if err != nil {
		return nil, err
	}
	x := &usageUIListEngineSpecsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UsageUI_ListEngineSpecsClient interface {
	Recv() (*ListEngineSpecsResponse, error)
	grpc.ClientStream
}

type usageUIListEngineSpecsClient struct {
	grpc.ClientStream
}

func (x *usageUIListEngineSpecsClient) Recv() (*ListEngineSpecsResponse, error) {
	m := new(ListEngineSpecsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *usageUIClient) IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error) {
	out := new(IsReadOnlyResponse)
	err := c.cc.Invoke(ctx, "/v1.UsageUI/IsReadOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsageUIServer is the server API for UsageUI service.
// All implementations must embed UnimplementedUsageUIServer
// for forward compatibility
type UsageUIServer interface {
	// ListEngineSpecs returns a list of Usage Engine(s) that can be started through the UI.
	ListEngineSpecs(*ListEngineSpecsRequest, UsageUI_ListEngineSpecsServer) error
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error)
	mustEmbedUnimplementedUsageUIServer()
}

// UnimplementedUsageUIServer must be embedded to have forward compatible implementations.
type UnimplementedUsageUIServer struct {
}

func (UnimplementedUsageUIServer) ListEngineSpecs(*ListEngineSpecsRequest, UsageUI_ListEngineSpecsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListEngineSpecs not implemented")
}
func (UnimplementedUsageUIServer) IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsReadOnly not implemented")
}
func (UnimplementedUsageUIServer) mustEmbedUnimplementedUsageUIServer() {}

// UnsafeUsageUIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsageUIServer will
// result in compilation errors.
type UnsafeUsageUIServer interface {
	mustEmbedUnimplementedUsageUIServer()
}

func RegisterUsageUIServer(s grpc.ServiceRegistrar, srv UsageUIServer) {
	s.RegisterService(&UsageUI_ServiceDesc, srv)
}

func _UsageUI_ListEngineSpecs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListEngineSpecsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UsageUIServer).ListEngineSpecs(m, &usageUIListEngineSpecsServer{stream})
}

type UsageUI_ListEngineSpecsServer interface {
	Send(*ListEngineSpecsResponse) error
	grpc.ServerStream
}

type usageUIListEngineSpecsServer struct {
	grpc.ServerStream
}

func (x *usageUIListEngineSpecsServer) Send(m *ListEngineSpecsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _UsageUI_IsReadOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsReadOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsageUIServer).IsReadOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UsageUI/IsReadOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsageUIServer).IsReadOnly(ctx, req.(*IsReadOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UsageUI_ServiceDesc is the grpc.ServiceDesc for UsageUI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UsageUI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.UsageUI",
	HandlerType: (*UsageUIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsReadOnly",
			Handler:    _UsageUI_IsReadOnly_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListEngineSpecs",
			Handler:       _UsageUI_ListEngineSpecs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "usage-ui.proto",
}
