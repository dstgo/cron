// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: api/cron/client/cron_client_service.proto

package v1

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

const (
	Task_Healthy_FullMethodName        = "/cron.api.cron.client.v1.Task/Healthy"
	Task_ExecTask_FullMethodName       = "/cron.api.cron.client.v1.Task/ExecTask"
	Task_CancelExecTask_FullMethodName = "/cron.api.cron.client.v1.Task/CancelExecTask"
)

// TaskClient is the client API for Task service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskClient interface {
	Healthy(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ExecTask(ctx context.Context, in *ExecTaskRequest, opts ...grpc.CallOption) (Task_ExecTaskClient, error)
	CancelExecTask(ctx context.Context, in *CancelExecTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type taskClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskClient(cc grpc.ClientConnInterface) TaskClient {
	return &taskClient{cc}
}

func (c *taskClient) Healthy(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Task_Healthy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) ExecTask(ctx context.Context, in *ExecTaskRequest, opts ...grpc.CallOption) (Task_ExecTaskClient, error) {
	stream, err := c.cc.NewStream(ctx, &Task_ServiceDesc.Streams[0], Task_ExecTask_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &taskExecTaskClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Task_ExecTaskClient interface {
	Recv() (*ExecTaskReply, error)
	grpc.ClientStream
}

type taskExecTaskClient struct {
	grpc.ClientStream
}

func (x *taskExecTaskClient) Recv() (*ExecTaskReply, error) {
	m := new(ExecTaskReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *taskClient) CancelExecTask(ctx context.Context, in *CancelExecTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Task_CancelExecTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServer is the server API for Task service.
// All implementations must embed UnimplementedTaskServer
// for forward compatibility
type TaskServer interface {
	Healthy(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	ExecTask(*ExecTaskRequest, Task_ExecTaskServer) error
	CancelExecTask(context.Context, *CancelExecTaskRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedTaskServer()
}

// UnimplementedTaskServer must be embedded to have forward compatible implementations.
type UnimplementedTaskServer struct {
}

func (UnimplementedTaskServer) Healthy(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Healthy not implemented")
}
func (UnimplementedTaskServer) ExecTask(*ExecTaskRequest, Task_ExecTaskServer) error {
	return status.Errorf(codes.Unimplemented, "method ExecTask not implemented")
}
func (UnimplementedTaskServer) CancelExecTask(context.Context, *CancelExecTaskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelExecTask not implemented")
}
func (UnimplementedTaskServer) mustEmbedUnimplementedTaskServer() {}

// UnsafeTaskServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskServer will
// result in compilation errors.
type UnsafeTaskServer interface {
	mustEmbedUnimplementedTaskServer()
}

func RegisterTaskServer(s grpc.ServiceRegistrar, srv TaskServer) {
	s.RegisterService(&Task_ServiceDesc, srv)
}

func _Task_Healthy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).Healthy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_Healthy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).Healthy(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_ExecTask_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ExecTaskRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TaskServer).ExecTask(m, &taskExecTaskServer{stream})
}

type Task_ExecTaskServer interface {
	Send(*ExecTaskReply) error
	grpc.ServerStream
}

type taskExecTaskServer struct {
	grpc.ServerStream
}

func (x *taskExecTaskServer) Send(m *ExecTaskReply) error {
	return x.ServerStream.SendMsg(m)
}

func _Task_CancelExecTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelExecTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).CancelExecTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_CancelExecTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).CancelExecTask(ctx, req.(*CancelExecTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Task_ServiceDesc is the grpc.ServiceDesc for Task service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Task_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cron.api.cron.client.v1.Task",
	HandlerType: (*TaskServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Healthy",
			Handler:    _Task_Healthy_Handler,
		},
		{
			MethodName: "CancelExecTask",
			Handler:    _Task_CancelExecTask_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ExecTask",
			Handler:       _Task_ExecTask_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/cron/client/cron_client_service.proto",
}
