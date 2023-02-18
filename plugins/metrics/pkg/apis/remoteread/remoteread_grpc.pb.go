// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - ragu               v1.0.0
// source: github.com/rancher/opni/plugins/metrics/pkg/apis/remoteread/remoteread.proto

package remoteread

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RemoteReadGatewayClient is the client API for RemoteReadGateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RemoteReadGatewayClient interface {
	AddTarget(ctx context.Context, in *TargetAddRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	EditTarget(ctx context.Context, in *TargetEditRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RemoveTarget(ctx context.Context, in *TargetRemoveRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListTargets(ctx context.Context, in *TargetListRequest, opts ...grpc.CallOption) (*TargetList, error)
	Start(ctx context.Context, in *StartReadRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Stop(ctx context.Context, in *StopReadRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetTargetStatus(ctx context.Context, in *TargetStatusRequest, opts ...grpc.CallOption) (*TargetStatus, error)
	Discover(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
}

type remoteReadGatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewRemoteReadGatewayClient(cc grpc.ClientConnInterface) RemoteReadGatewayClient {
	return &remoteReadGatewayClient{cc}
}

func (c *remoteReadGatewayClient) AddTarget(ctx context.Context, in *TargetAddRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/remoteread.RemoteReadGateway/AddTarget", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteReadGatewayClient) EditTarget(ctx context.Context, in *TargetEditRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/remoteread.RemoteReadGateway/EditTarget", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteReadGatewayClient) RemoveTarget(ctx context.Context, in *TargetRemoveRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/remoteread.RemoteReadGateway/RemoveTarget", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteReadGatewayClient) ListTargets(ctx context.Context, in *TargetListRequest, opts ...grpc.CallOption) (*TargetList, error) {
	out := new(TargetList)
	err := c.cc.Invoke(ctx, "/remoteread.RemoteReadGateway/ListTargets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteReadGatewayClient) Start(ctx context.Context, in *StartReadRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/remoteread.RemoteReadGateway/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteReadGatewayClient) Stop(ctx context.Context, in *StopReadRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/remoteread.RemoteReadGateway/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteReadGatewayClient) GetTargetStatus(ctx context.Context, in *TargetStatusRequest, opts ...grpc.CallOption) (*TargetStatus, error) {
	out := new(TargetStatus)
	err := c.cc.Invoke(ctx, "/remoteread.RemoteReadGateway/GetTargetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteReadGatewayClient) Discover(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/remoteread.RemoteReadGateway/Discover", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemoteReadGatewayServer is the server API for RemoteReadGateway service.
// All implementations must embed UnimplementedRemoteReadGatewayServer
// for forward compatibility
type RemoteReadGatewayServer interface {
	AddTarget(context.Context, *TargetAddRequest) (*emptypb.Empty, error)
	EditTarget(context.Context, *TargetEditRequest) (*emptypb.Empty, error)
	RemoveTarget(context.Context, *TargetRemoveRequest) (*emptypb.Empty, error)
	ListTargets(context.Context, *TargetListRequest) (*TargetList, error)
	Start(context.Context, *StartReadRequest) (*emptypb.Empty, error)
	Stop(context.Context, *StopReadRequest) (*emptypb.Empty, error)
	GetTargetStatus(context.Context, *TargetStatusRequest) (*TargetStatus, error)
	Discover(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error)
	mustEmbedUnimplementedRemoteReadGatewayServer()
}

// UnimplementedRemoteReadGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedRemoteReadGatewayServer struct {
}

func (UnimplementedRemoteReadGatewayServer) AddTarget(context.Context, *TargetAddRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTarget not implemented")
}
func (UnimplementedRemoteReadGatewayServer) EditTarget(context.Context, *TargetEditRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditTarget not implemented")
}
func (UnimplementedRemoteReadGatewayServer) RemoveTarget(context.Context, *TargetRemoveRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveTarget not implemented")
}
func (UnimplementedRemoteReadGatewayServer) ListTargets(context.Context, *TargetListRequest) (*TargetList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTargets not implemented")
}
func (UnimplementedRemoteReadGatewayServer) Start(context.Context, *StartReadRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedRemoteReadGatewayServer) Stop(context.Context, *StopReadRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedRemoteReadGatewayServer) GetTargetStatus(context.Context, *TargetStatusRequest) (*TargetStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTargetStatus not implemented")
}
func (UnimplementedRemoteReadGatewayServer) Discover(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Discover not implemented")
}
func (UnimplementedRemoteReadGatewayServer) mustEmbedUnimplementedRemoteReadGatewayServer() {}

// UnsafeRemoteReadGatewayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RemoteReadGatewayServer will
// result in compilation errors.
type UnsafeRemoteReadGatewayServer interface {
	mustEmbedUnimplementedRemoteReadGatewayServer()
}

func RegisterRemoteReadGatewayServer(s grpc.ServiceRegistrar, srv RemoteReadGatewayServer) {
	s.RegisterService(&RemoteReadGateway_ServiceDesc, srv)
}

func _RemoteReadGateway_AddTarget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TargetAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteReadGatewayServer).AddTarget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remoteread.RemoteReadGateway/AddTarget",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteReadGatewayServer).AddTarget(ctx, req.(*TargetAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteReadGateway_EditTarget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TargetEditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteReadGatewayServer).EditTarget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remoteread.RemoteReadGateway/EditTarget",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteReadGatewayServer).EditTarget(ctx, req.(*TargetEditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteReadGateway_RemoveTarget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TargetRemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteReadGatewayServer).RemoveTarget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remoteread.RemoteReadGateway/RemoveTarget",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteReadGatewayServer).RemoveTarget(ctx, req.(*TargetRemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteReadGateway_ListTargets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TargetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteReadGatewayServer).ListTargets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remoteread.RemoteReadGateway/ListTargets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteReadGatewayServer).ListTargets(ctx, req.(*TargetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteReadGateway_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteReadGatewayServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remoteread.RemoteReadGateway/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteReadGatewayServer).Start(ctx, req.(*StartReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteReadGateway_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteReadGatewayServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remoteread.RemoteReadGateway/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteReadGatewayServer).Stop(ctx, req.(*StopReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteReadGateway_GetTargetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TargetStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteReadGatewayServer).GetTargetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remoteread.RemoteReadGateway/GetTargetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteReadGatewayServer).GetTargetStatus(ctx, req.(*TargetStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteReadGateway_Discover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteReadGatewayServer).Discover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remoteread.RemoteReadGateway/Discover",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteReadGatewayServer).Discover(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RemoteReadGateway_ServiceDesc is the grpc.ServiceDesc for RemoteReadGateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RemoteReadGateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "remoteread.RemoteReadGateway",
	HandlerType: (*RemoteReadGatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTarget",
			Handler:    _RemoteReadGateway_AddTarget_Handler,
		},
		{
			MethodName: "EditTarget",
			Handler:    _RemoteReadGateway_EditTarget_Handler,
		},
		{
			MethodName: "RemoveTarget",
			Handler:    _RemoteReadGateway_RemoveTarget_Handler,
		},
		{
			MethodName: "ListTargets",
			Handler:    _RemoteReadGateway_ListTargets_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _RemoteReadGateway_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _RemoteReadGateway_Stop_Handler,
		},
		{
			MethodName: "GetTargetStatus",
			Handler:    _RemoteReadGateway_GetTargetStatus_Handler,
		},
		{
			MethodName: "Discover",
			Handler:    _RemoteReadGateway_Discover_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rancher/opni/plugins/metrics/pkg/apis/remoteread/remoteread.proto",
}

// RemoteReadAgentClient is the client API for RemoteReadAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RemoteReadAgentClient interface {
	Start(ctx context.Context, in *StartReadRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Stop(ctx context.Context, in *StopReadRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetTargetStatus(ctx context.Context, in *TargetStatusRequest, opts ...grpc.CallOption) (*TargetStatus, error)
	Discover(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
}

type remoteReadAgentClient struct {
	cc grpc.ClientConnInterface
}

func NewRemoteReadAgentClient(cc grpc.ClientConnInterface) RemoteReadAgentClient {
	return &remoteReadAgentClient{cc}
}

func (c *remoteReadAgentClient) Start(ctx context.Context, in *StartReadRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/remoteread.RemoteReadAgent/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteReadAgentClient) Stop(ctx context.Context, in *StopReadRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/remoteread.RemoteReadAgent/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteReadAgentClient) GetTargetStatus(ctx context.Context, in *TargetStatusRequest, opts ...grpc.CallOption) (*TargetStatus, error) {
	out := new(TargetStatus)
	err := c.cc.Invoke(ctx, "/remoteread.RemoteReadAgent/GetTargetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteReadAgentClient) Discover(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/remoteread.RemoteReadAgent/Discover", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemoteReadAgentServer is the server API for RemoteReadAgent service.
// All implementations must embed UnimplementedRemoteReadAgentServer
// for forward compatibility
type RemoteReadAgentServer interface {
	Start(context.Context, *StartReadRequest) (*emptypb.Empty, error)
	Stop(context.Context, *StopReadRequest) (*emptypb.Empty, error)
	GetTargetStatus(context.Context, *TargetStatusRequest) (*TargetStatus, error)
	Discover(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error)
	mustEmbedUnimplementedRemoteReadAgentServer()
}

// UnimplementedRemoteReadAgentServer must be embedded to have forward compatible implementations.
type UnimplementedRemoteReadAgentServer struct {
}

func (UnimplementedRemoteReadAgentServer) Start(context.Context, *StartReadRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedRemoteReadAgentServer) Stop(context.Context, *StopReadRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedRemoteReadAgentServer) GetTargetStatus(context.Context, *TargetStatusRequest) (*TargetStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTargetStatus not implemented")
}
func (UnimplementedRemoteReadAgentServer) Discover(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Discover not implemented")
}
func (UnimplementedRemoteReadAgentServer) mustEmbedUnimplementedRemoteReadAgentServer() {}

// UnsafeRemoteReadAgentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RemoteReadAgentServer will
// result in compilation errors.
type UnsafeRemoteReadAgentServer interface {
	mustEmbedUnimplementedRemoteReadAgentServer()
}

func RegisterRemoteReadAgentServer(s grpc.ServiceRegistrar, srv RemoteReadAgentServer) {
	s.RegisterService(&RemoteReadAgent_ServiceDesc, srv)
}

func _RemoteReadAgent_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteReadAgentServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remoteread.RemoteReadAgent/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteReadAgentServer).Start(ctx, req.(*StartReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteReadAgent_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteReadAgentServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remoteread.RemoteReadAgent/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteReadAgentServer).Stop(ctx, req.(*StopReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteReadAgent_GetTargetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TargetStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteReadAgentServer).GetTargetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remoteread.RemoteReadAgent/GetTargetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteReadAgentServer).GetTargetStatus(ctx, req.(*TargetStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteReadAgent_Discover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteReadAgentServer).Discover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remoteread.RemoteReadAgent/Discover",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteReadAgentServer).Discover(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RemoteReadAgent_ServiceDesc is the grpc.ServiceDesc for RemoteReadAgent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RemoteReadAgent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "remoteread.RemoteReadAgent",
	HandlerType: (*RemoteReadAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _RemoteReadAgent_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _RemoteReadAgent_Stop_Handler,
		},
		{
			MethodName: "GetTargetStatus",
			Handler:    _RemoteReadAgent_GetTargetStatus_Handler,
		},
		{
			MethodName: "Discover",
			Handler:    _RemoteReadAgent_Discover_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rancher/opni/plugins/metrics/pkg/apis/remoteread/remoteread.proto",
}
