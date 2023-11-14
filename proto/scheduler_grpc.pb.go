// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.0
// source: proto/scheduler.proto

package proto

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

// ScheduleServiceClient is the client API for ScheduleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScheduleServiceClient interface {
	CreateSchedule(ctx context.Context, in *CreateScheduleRequest, opts ...grpc.CallOption) (*CreateScheduleResponse, error)
	SendError(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*SendErrorResponse, error)
}

type scheduleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewScheduleServiceClient(cc grpc.ClientConnInterface) ScheduleServiceClient {
	return &scheduleServiceClient{cc}
}

func (c *scheduleServiceClient) CreateSchedule(ctx context.Context, in *CreateScheduleRequest, opts ...grpc.CallOption) (*CreateScheduleResponse, error) {
	out := new(CreateScheduleResponse)
	err := c.cc.Invoke(ctx, "/schedule_service.ScheduleService/CreateSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleServiceClient) SendError(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*SendErrorResponse, error) {
	out := new(SendErrorResponse)
	err := c.cc.Invoke(ctx, "/schedule_service.ScheduleService/SendError", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScheduleServiceServer is the server API for ScheduleService service.
// All implementations must embed UnimplementedScheduleServiceServer
// for forward compatibility
type ScheduleServiceServer interface {
	CreateSchedule(context.Context, *CreateScheduleRequest) (*CreateScheduleResponse, error)
	SendError(context.Context, *NoParam) (*SendErrorResponse, error)
	mustEmbedUnimplementedScheduleServiceServer()
}

// UnimplementedScheduleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedScheduleServiceServer struct {
}

func (UnimplementedScheduleServiceServer) CreateSchedule(context.Context, *CreateScheduleRequest) (*CreateScheduleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSchedule not implemented")
}
func (UnimplementedScheduleServiceServer) SendError(context.Context, *NoParam) (*SendErrorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendError not implemented")
}
func (UnimplementedScheduleServiceServer) mustEmbedUnimplementedScheduleServiceServer() {}

// UnsafeScheduleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScheduleServiceServer will
// result in compilation errors.
type UnsafeScheduleServiceServer interface {
	mustEmbedUnimplementedScheduleServiceServer()
}

func RegisterScheduleServiceServer(s grpc.ServiceRegistrar, srv ScheduleServiceServer) {
	s.RegisterService(&ScheduleService_ServiceDesc, srv)
}

func _ScheduleService_CreateSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServiceServer).CreateSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schedule_service.ScheduleService/CreateSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServiceServer).CreateSchedule(ctx, req.(*CreateScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScheduleService_SendError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServiceServer).SendError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schedule_service.ScheduleService/SendError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServiceServer).SendError(ctx, req.(*NoParam))
	}
	return interceptor(ctx, in, info, handler)
}

// ScheduleService_ServiceDesc is the grpc.ServiceDesc for ScheduleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ScheduleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "schedule_service.ScheduleService",
	HandlerType: (*ScheduleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSchedule",
			Handler:    _ScheduleService_CreateSchedule_Handler,
		},
		{
			MethodName: "SendError",
			Handler:    _ScheduleService_SendError_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/scheduler.proto",
}
