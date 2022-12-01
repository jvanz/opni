// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - ragu               v1.0.0
// source: github.com/rancher/opni/pkg/apis/control/v1/remote.proto

package v1

import (
	context "context"
	v1 "github.com/rancher/opni/pkg/apis/core/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HealthClient is the client API for Health service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HealthClient interface {
	GetHealth(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.Health, error)
}

type healthClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthClient(cc grpc.ClientConnInterface) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) GetHealth(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.Health, error) {
	out := new(v1.Health)
	err := c.cc.Invoke(ctx, "/control.Health/GetHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthServer is the server API for Health service.
// All implementations must embed UnimplementedHealthServer
// for forward compatibility
type HealthServer interface {
	GetHealth(context.Context, *emptypb.Empty) (*v1.Health, error)
	mustEmbedUnimplementedHealthServer()
}

// UnimplementedHealthServer must be embedded to have forward compatible implementations.
type UnimplementedHealthServer struct {
}

func (UnimplementedHealthServer) GetHealth(context.Context, *emptypb.Empty) (*v1.Health, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHealth not implemented")
}
func (UnimplementedHealthServer) mustEmbedUnimplementedHealthServer() {}

// UnsafeHealthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthServer will
// result in compilation errors.
type UnsafeHealthServer interface {
	mustEmbedUnimplementedHealthServer()
}

func RegisterHealthServer(s grpc.ServiceRegistrar, srv HealthServer) {
	s.RegisterService(&Health_ServiceDesc, srv)
}

func _Health_GetHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).GetHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control.Health/GetHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).GetHealth(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Health_ServiceDesc is the grpc.ServiceDesc for Health service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Health_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "control.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHealth",
			Handler:    _Health_GetHealth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rancher/opni/pkg/apis/control/v1/remote.proto",
}

// HealthListenerClient is the client API for HealthListener service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HealthListenerClient interface {
	UpdateHealth(ctx context.Context, in *v1.Health, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type healthListenerClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthListenerClient(cc grpc.ClientConnInterface) HealthListenerClient {
	return &healthListenerClient{cc}
}

func (c *healthListenerClient) UpdateHealth(ctx context.Context, in *v1.Health, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/control.HealthListener/UpdateHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthListenerServer is the server API for HealthListener service.
// All implementations must embed UnimplementedHealthListenerServer
// for forward compatibility
type HealthListenerServer interface {
	UpdateHealth(context.Context, *v1.Health) (*emptypb.Empty, error)
	mustEmbedUnimplementedHealthListenerServer()
}

// UnimplementedHealthListenerServer must be embedded to have forward compatible implementations.
type UnimplementedHealthListenerServer struct {
}

func (UnimplementedHealthListenerServer) UpdateHealth(context.Context, *v1.Health) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHealth not implemented")
}
func (UnimplementedHealthListenerServer) mustEmbedUnimplementedHealthListenerServer() {}

// UnsafeHealthListenerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthListenerServer will
// result in compilation errors.
type UnsafeHealthListenerServer interface {
	mustEmbedUnimplementedHealthListenerServer()
}

func RegisterHealthListenerServer(s grpc.ServiceRegistrar, srv HealthListenerServer) {
	s.RegisterService(&HealthListener_ServiceDesc, srv)
}

func _HealthListener_UpdateHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.Health)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthListenerServer).UpdateHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control.HealthListener/UpdateHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthListenerServer).UpdateHealth(ctx, req.(*v1.Health))
	}
	return interceptor(ctx, in, info, handler)
}

// HealthListener_ServiceDesc is the grpc.ServiceDesc for HealthListener service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HealthListener_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "control.HealthListener",
	HandlerType: (*HealthListenerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateHealth",
			Handler:    _HealthListener_UpdateHealth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rancher/opni/pkg/apis/control/v1/remote.proto",
}

// PluginSyncClient is the client API for PluginSync service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PluginSyncClient interface {
	SyncPluginManifest(ctx context.Context, in *PluginManifest, opts ...grpc.CallOption) (*PluginArchive, error)
	GetPluginManifest(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PluginManifest, error)
}

type pluginSyncClient struct {
	cc grpc.ClientConnInterface
}

func NewPluginSyncClient(cc grpc.ClientConnInterface) PluginSyncClient {
	return &pluginSyncClient{cc}
}

func (c *pluginSyncClient) SyncPluginManifest(ctx context.Context, in *PluginManifest, opts ...grpc.CallOption) (*PluginArchive, error) {
	out := new(PluginArchive)
	err := c.cc.Invoke(ctx, "/control.PluginSync/SyncPluginManifest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginSyncClient) GetPluginManifest(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PluginManifest, error) {
	out := new(PluginManifest)
	err := c.cc.Invoke(ctx, "/control.PluginSync/GetPluginManifest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PluginSyncServer is the server API for PluginSync service.
// All implementations must embed UnimplementedPluginSyncServer
// for forward compatibility
type PluginSyncServer interface {
	SyncPluginManifest(context.Context, *PluginManifest) (*PluginArchive, error)
	GetPluginManifest(context.Context, *emptypb.Empty) (*PluginManifest, error)
	mustEmbedUnimplementedPluginSyncServer()
}

// UnimplementedPluginSyncServer must be embedded to have forward compatible implementations.
type UnimplementedPluginSyncServer struct {
}

func (UnimplementedPluginSyncServer) SyncPluginManifest(context.Context, *PluginManifest) (*PluginArchive, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncPluginManifest not implemented")
}
func (UnimplementedPluginSyncServer) GetPluginManifest(context.Context, *emptypb.Empty) (*PluginManifest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPluginManifest not implemented")
}
func (UnimplementedPluginSyncServer) mustEmbedUnimplementedPluginSyncServer() {}

// UnsafePluginSyncServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PluginSyncServer will
// result in compilation errors.
type UnsafePluginSyncServer interface {
	mustEmbedUnimplementedPluginSyncServer()
}

func RegisterPluginSyncServer(s grpc.ServiceRegistrar, srv PluginSyncServer) {
	s.RegisterService(&PluginSync_ServiceDesc, srv)
}

func _PluginSync_SyncPluginManifest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginManifest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginSyncServer).SyncPluginManifest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control.PluginSync/SyncPluginManifest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginSyncServer).SyncPluginManifest(ctx, req.(*PluginManifest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PluginSync_GetPluginManifest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginSyncServer).GetPluginManifest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control.PluginSync/GetPluginManifest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginSyncServer).GetPluginManifest(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// PluginSync_ServiceDesc is the grpc.ServiceDesc for PluginSync service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PluginSync_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "control.PluginSync",
	HandlerType: (*PluginSyncServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SyncPluginManifest",
			Handler:    _PluginSync_SyncPluginManifest_Handler,
		},
		{
			MethodName: "GetPluginManifest",
			Handler:    _PluginSync_GetPluginManifest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rancher/opni/pkg/apis/control/v1/remote.proto",
}
