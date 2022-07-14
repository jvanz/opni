// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - ragu               v1.0.0
// source: github.com/rancher/opni/pkg/plugins/alerting/pkg/apis/alerting/alerting.proto

package alerting

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

// AlertingClient is the client API for Alerting service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AlertingClient interface {
	// ------- Trigger Alerts -------
	// TODO
	TriggerAlerts(ctx context.Context, in *TriggerAlertsRequest, opts ...grpc.CallOption) (*TriggerAlertsResponse, error)
	// ------- CRUD for alerting events log -------
	CreateAlertLog(ctx context.Context, in *AlertLog, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// id is the unix epoch timestamp of the alert
	GetAlertLog(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListAlertLogs(ctx context.Context, in *ListAlertLogRequest, opts ...grpc.CallOption) (*AlertLogList, error)
	UpdateAlertLog(ctx context.Context, in *UpdateAlertLogRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteAlertLog(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// -------- CRUD for alerting conditions -------
	//TODO
	CreateAlertCondition(ctx context.Context, in *AlertCondition, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetAlertCondition(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*AlertCondition, error)
	ListAlertConditions(ctx context.Context, in *ListAlertConditionRequest, opts ...grpc.CallOption) (*AlertConditionList, error)
	UpdateAlertCondition(ctx context.Context, in *UpdateAlertConditionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteAlertCondition(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*emptypb.Empty, error)
	PreviewAlertCondition(ctx context.Context, in *PreviewAlertConditionRequest, opts ...grpc.CallOption) (*PreviewAlertConditionResponse, error)
	// -------- CRUD for alerting endpoints ----
	// TODO
	CreateAlertEndpoint(ctx context.Context, in *AlertEndpoint, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetAlertEndpoint(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*AlertEndpoint, error)
	ListAlertEndpoints(ctx context.Context, in *ListAlertEndpointsRequest, opts ...grpc.CallOption) (*AlertEndpointList, error)
	UpdateAlertEndpoint(ctx context.Context, in *UpdateAlertEndpointRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteAlertEndpoint(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*emptypb.Empty, error)
	TestAlertEndpoint(ctx context.Context, in *TestAlertEndpointRequest, opts ...grpc.CallOption) (*TestAlertEndpointResponse, error)
}

type alertingClient struct {
	cc grpc.ClientConnInterface
}

func NewAlertingClient(cc grpc.ClientConnInterface) AlertingClient {
	return &alertingClient{cc}
}

func (c *alertingClient) TriggerAlerts(ctx context.Context, in *TriggerAlertsRequest, opts ...grpc.CallOption) (*TriggerAlertsResponse, error) {
	out := new(TriggerAlertsResponse)
	err := c.cc.Invoke(ctx, "/Alerting/TriggerAlerts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertingClient) CreateAlertLog(ctx context.Context, in *AlertLog, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Alerting/CreateAlertLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertingClient) GetAlertLog(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Alerting/GetAlertLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertingClient) ListAlertLogs(ctx context.Context, in *ListAlertLogRequest, opts ...grpc.CallOption) (*AlertLogList, error) {
	out := new(AlertLogList)
	err := c.cc.Invoke(ctx, "/Alerting/ListAlertLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertingClient) UpdateAlertLog(ctx context.Context, in *UpdateAlertLogRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Alerting/UpdateAlertLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertingClient) DeleteAlertLog(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Alerting/DeleteAlertLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertingClient) CreateAlertCondition(ctx context.Context, in *AlertCondition, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Alerting/CreateAlertCondition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertingClient) GetAlertCondition(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*AlertCondition, error) {
	out := new(AlertCondition)
	err := c.cc.Invoke(ctx, "/Alerting/GetAlertCondition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertingClient) ListAlertConditions(ctx context.Context, in *ListAlertConditionRequest, opts ...grpc.CallOption) (*AlertConditionList, error) {
	out := new(AlertConditionList)
	err := c.cc.Invoke(ctx, "/Alerting/ListAlertConditions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertingClient) UpdateAlertCondition(ctx context.Context, in *UpdateAlertConditionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Alerting/UpdateAlertCondition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertingClient) DeleteAlertCondition(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Alerting/DeleteAlertCondition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertingClient) PreviewAlertCondition(ctx context.Context, in *PreviewAlertConditionRequest, opts ...grpc.CallOption) (*PreviewAlertConditionResponse, error) {
	out := new(PreviewAlertConditionResponse)
	err := c.cc.Invoke(ctx, "/Alerting/PreviewAlertCondition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertingClient) CreateAlertEndpoint(ctx context.Context, in *AlertEndpoint, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Alerting/CreateAlertEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertingClient) GetAlertEndpoint(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*AlertEndpoint, error) {
	out := new(AlertEndpoint)
	err := c.cc.Invoke(ctx, "/Alerting/GetAlertEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertingClient) ListAlertEndpoints(ctx context.Context, in *ListAlertEndpointsRequest, opts ...grpc.CallOption) (*AlertEndpointList, error) {
	out := new(AlertEndpointList)
	err := c.cc.Invoke(ctx, "/Alerting/ListAlertEndpoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertingClient) UpdateAlertEndpoint(ctx context.Context, in *UpdateAlertEndpointRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Alerting/UpdateAlertEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertingClient) DeleteAlertEndpoint(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Alerting/DeleteAlertEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertingClient) TestAlertEndpoint(ctx context.Context, in *TestAlertEndpointRequest, opts ...grpc.CallOption) (*TestAlertEndpointResponse, error) {
	out := new(TestAlertEndpointResponse)
	err := c.cc.Invoke(ctx, "/Alerting/TestAlertEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlertingServer is the server API for Alerting service.
// All implementations must embed UnimplementedAlertingServer
// for forward compatibility
type AlertingServer interface {
	// ------- Trigger Alerts -------
	// TODO
	TriggerAlerts(context.Context, *TriggerAlertsRequest) (*TriggerAlertsResponse, error)
	// ------- CRUD for alerting events log -------
	CreateAlertLog(context.Context, *AlertLog) (*emptypb.Empty, error)
	// id is the unix epoch timestamp of the alert
	GetAlertLog(context.Context, *v1.Reference) (*emptypb.Empty, error)
	ListAlertLogs(context.Context, *ListAlertLogRequest) (*AlertLogList, error)
	UpdateAlertLog(context.Context, *UpdateAlertLogRequest) (*emptypb.Empty, error)
	DeleteAlertLog(context.Context, *v1.Reference) (*emptypb.Empty, error)
	// -------- CRUD for alerting conditions -------
	//TODO
	CreateAlertCondition(context.Context, *AlertCondition) (*emptypb.Empty, error)
	GetAlertCondition(context.Context, *v1.Reference) (*AlertCondition, error)
	ListAlertConditions(context.Context, *ListAlertConditionRequest) (*AlertConditionList, error)
	UpdateAlertCondition(context.Context, *UpdateAlertConditionRequest) (*emptypb.Empty, error)
	DeleteAlertCondition(context.Context, *v1.Reference) (*emptypb.Empty, error)
	PreviewAlertCondition(context.Context, *PreviewAlertConditionRequest) (*PreviewAlertConditionResponse, error)
	// -------- CRUD for alerting endpoints ----
	// TODO
	CreateAlertEndpoint(context.Context, *AlertEndpoint) (*emptypb.Empty, error)
	GetAlertEndpoint(context.Context, *v1.Reference) (*AlertEndpoint, error)
	ListAlertEndpoints(context.Context, *ListAlertEndpointsRequest) (*AlertEndpointList, error)
	UpdateAlertEndpoint(context.Context, *UpdateAlertEndpointRequest) (*emptypb.Empty, error)
	DeleteAlertEndpoint(context.Context, *v1.Reference) (*emptypb.Empty, error)
	TestAlertEndpoint(context.Context, *TestAlertEndpointRequest) (*TestAlertEndpointResponse, error)
	mustEmbedUnimplementedAlertingServer()
}

// UnimplementedAlertingServer must be embedded to have forward compatible implementations.
type UnimplementedAlertingServer struct {
}

func (UnimplementedAlertingServer) TriggerAlerts(context.Context, *TriggerAlertsRequest) (*TriggerAlertsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TriggerAlerts not implemented")
}
func (UnimplementedAlertingServer) CreateAlertLog(context.Context, *AlertLog) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAlertLog not implemented")
}
func (UnimplementedAlertingServer) GetAlertLog(context.Context, *v1.Reference) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlertLog not implemented")
}
func (UnimplementedAlertingServer) ListAlertLogs(context.Context, *ListAlertLogRequest) (*AlertLogList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAlertLogs not implemented")
}
func (UnimplementedAlertingServer) UpdateAlertLog(context.Context, *UpdateAlertLogRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAlertLog not implemented")
}
func (UnimplementedAlertingServer) DeleteAlertLog(context.Context, *v1.Reference) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAlertLog not implemented")
}
func (UnimplementedAlertingServer) CreateAlertCondition(context.Context, *AlertCondition) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAlertCondition not implemented")
}
func (UnimplementedAlertingServer) GetAlertCondition(context.Context, *v1.Reference) (*AlertCondition, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlertCondition not implemented")
}
func (UnimplementedAlertingServer) ListAlertConditions(context.Context, *ListAlertConditionRequest) (*AlertConditionList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAlertConditions not implemented")
}
func (UnimplementedAlertingServer) UpdateAlertCondition(context.Context, *UpdateAlertConditionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAlertCondition not implemented")
}
func (UnimplementedAlertingServer) DeleteAlertCondition(context.Context, *v1.Reference) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAlertCondition not implemented")
}
func (UnimplementedAlertingServer) PreviewAlertCondition(context.Context, *PreviewAlertConditionRequest) (*PreviewAlertConditionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewAlertCondition not implemented")
}
func (UnimplementedAlertingServer) CreateAlertEndpoint(context.Context, *AlertEndpoint) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAlertEndpoint not implemented")
}
func (UnimplementedAlertingServer) GetAlertEndpoint(context.Context, *v1.Reference) (*AlertEndpoint, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlertEndpoint not implemented")
}
func (UnimplementedAlertingServer) ListAlertEndpoints(context.Context, *ListAlertEndpointsRequest) (*AlertEndpointList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAlertEndpoints not implemented")
}
func (UnimplementedAlertingServer) UpdateAlertEndpoint(context.Context, *UpdateAlertEndpointRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAlertEndpoint not implemented")
}
func (UnimplementedAlertingServer) DeleteAlertEndpoint(context.Context, *v1.Reference) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAlertEndpoint not implemented")
}
func (UnimplementedAlertingServer) TestAlertEndpoint(context.Context, *TestAlertEndpointRequest) (*TestAlertEndpointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestAlertEndpoint not implemented")
}
func (UnimplementedAlertingServer) mustEmbedUnimplementedAlertingServer() {}

// UnsafeAlertingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AlertingServer will
// result in compilation errors.
type UnsafeAlertingServer interface {
	mustEmbedUnimplementedAlertingServer()
}

func RegisterAlertingServer(s grpc.ServiceRegistrar, srv AlertingServer) {
	s.RegisterService(&Alerting_ServiceDesc, srv)
}

func _Alerting_TriggerAlerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TriggerAlertsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertingServer).TriggerAlerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Alerting/TriggerAlerts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertingServer).TriggerAlerts(ctx, req.(*TriggerAlertsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerting_CreateAlertLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertLog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertingServer).CreateAlertLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Alerting/CreateAlertLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertingServer).CreateAlertLog(ctx, req.(*AlertLog))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerting_GetAlertLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.Reference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertingServer).GetAlertLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Alerting/GetAlertLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertingServer).GetAlertLog(ctx, req.(*v1.Reference))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerting_ListAlertLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAlertLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertingServer).ListAlertLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Alerting/ListAlertLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertingServer).ListAlertLogs(ctx, req.(*ListAlertLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerting_UpdateAlertLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAlertLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertingServer).UpdateAlertLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Alerting/UpdateAlertLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertingServer).UpdateAlertLog(ctx, req.(*UpdateAlertLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerting_DeleteAlertLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.Reference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertingServer).DeleteAlertLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Alerting/DeleteAlertLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertingServer).DeleteAlertLog(ctx, req.(*v1.Reference))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerting_CreateAlertCondition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertCondition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertingServer).CreateAlertCondition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Alerting/CreateAlertCondition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertingServer).CreateAlertCondition(ctx, req.(*AlertCondition))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerting_GetAlertCondition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.Reference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertingServer).GetAlertCondition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Alerting/GetAlertCondition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertingServer).GetAlertCondition(ctx, req.(*v1.Reference))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerting_ListAlertConditions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAlertConditionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertingServer).ListAlertConditions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Alerting/ListAlertConditions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertingServer).ListAlertConditions(ctx, req.(*ListAlertConditionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerting_UpdateAlertCondition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAlertConditionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertingServer).UpdateAlertCondition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Alerting/UpdateAlertCondition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertingServer).UpdateAlertCondition(ctx, req.(*UpdateAlertConditionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerting_DeleteAlertCondition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.Reference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertingServer).DeleteAlertCondition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Alerting/DeleteAlertCondition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertingServer).DeleteAlertCondition(ctx, req.(*v1.Reference))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerting_PreviewAlertCondition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PreviewAlertConditionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertingServer).PreviewAlertCondition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Alerting/PreviewAlertCondition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertingServer).PreviewAlertCondition(ctx, req.(*PreviewAlertConditionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerting_CreateAlertEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertEndpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertingServer).CreateAlertEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Alerting/CreateAlertEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertingServer).CreateAlertEndpoint(ctx, req.(*AlertEndpoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerting_GetAlertEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.Reference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertingServer).GetAlertEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Alerting/GetAlertEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertingServer).GetAlertEndpoint(ctx, req.(*v1.Reference))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerting_ListAlertEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAlertEndpointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertingServer).ListAlertEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Alerting/ListAlertEndpoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertingServer).ListAlertEndpoints(ctx, req.(*ListAlertEndpointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerting_UpdateAlertEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAlertEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertingServer).UpdateAlertEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Alerting/UpdateAlertEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertingServer).UpdateAlertEndpoint(ctx, req.(*UpdateAlertEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerting_DeleteAlertEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.Reference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertingServer).DeleteAlertEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Alerting/DeleteAlertEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertingServer).DeleteAlertEndpoint(ctx, req.(*v1.Reference))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerting_TestAlertEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestAlertEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertingServer).TestAlertEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Alerting/TestAlertEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertingServer).TestAlertEndpoint(ctx, req.(*TestAlertEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Alerting_ServiceDesc is the grpc.ServiceDesc for Alerting service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Alerting_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Alerting",
	HandlerType: (*AlertingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TriggerAlerts",
			Handler:    _Alerting_TriggerAlerts_Handler,
		},
		{
			MethodName: "CreateAlertLog",
			Handler:    _Alerting_CreateAlertLog_Handler,
		},
		{
			MethodName: "GetAlertLog",
			Handler:    _Alerting_GetAlertLog_Handler,
		},
		{
			MethodName: "ListAlertLogs",
			Handler:    _Alerting_ListAlertLogs_Handler,
		},
		{
			MethodName: "UpdateAlertLog",
			Handler:    _Alerting_UpdateAlertLog_Handler,
		},
		{
			MethodName: "DeleteAlertLog",
			Handler:    _Alerting_DeleteAlertLog_Handler,
		},
		{
			MethodName: "CreateAlertCondition",
			Handler:    _Alerting_CreateAlertCondition_Handler,
		},
		{
			MethodName: "GetAlertCondition",
			Handler:    _Alerting_GetAlertCondition_Handler,
		},
		{
			MethodName: "ListAlertConditions",
			Handler:    _Alerting_ListAlertConditions_Handler,
		},
		{
			MethodName: "UpdateAlertCondition",
			Handler:    _Alerting_UpdateAlertCondition_Handler,
		},
		{
			MethodName: "DeleteAlertCondition",
			Handler:    _Alerting_DeleteAlertCondition_Handler,
		},
		{
			MethodName: "PreviewAlertCondition",
			Handler:    _Alerting_PreviewAlertCondition_Handler,
		},
		{
			MethodName: "CreateAlertEndpoint",
			Handler:    _Alerting_CreateAlertEndpoint_Handler,
		},
		{
			MethodName: "GetAlertEndpoint",
			Handler:    _Alerting_GetAlertEndpoint_Handler,
		},
		{
			MethodName: "ListAlertEndpoints",
			Handler:    _Alerting_ListAlertEndpoints_Handler,
		},
		{
			MethodName: "UpdateAlertEndpoint",
			Handler:    _Alerting_UpdateAlertEndpoint_Handler,
		},
		{
			MethodName: "DeleteAlertEndpoint",
			Handler:    _Alerting_DeleteAlertEndpoint_Handler,
		},
		{
			MethodName: "TestAlertEndpoint",
			Handler:    _Alerting_TestAlertEndpoint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rancher/opni/pkg/plugins/alerting/pkg/apis/alerting/alerting.proto",
}
