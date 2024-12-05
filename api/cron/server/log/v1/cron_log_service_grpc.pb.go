// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: api/cron/server/log/cron_log_service.proto

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

const (
	Log_GetLog_FullMethodName  = "/cron.api.server.cron.log.v1.Log/GetLog"
	Log_ListLog_FullMethodName = "/cron.api.server.cron.log.v1.Log/ListLog"
)

// LogClient is the client API for Log service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogClient interface {
	// GetLog 获取指定的日志信息
	GetLog(ctx context.Context, in *GetLogRequest, opts ...grpc.CallOption) (*GetLogReply, error)
	// ListLog 获取日志信息列表
	ListLog(ctx context.Context, in *ListLogRequest, opts ...grpc.CallOption) (*ListLogReply, error)
}

type logClient struct {
	cc grpc.ClientConnInterface
}

func NewLogClient(cc grpc.ClientConnInterface) LogClient {
	return &logClient{cc}
}

func (c *logClient) GetLog(ctx context.Context, in *GetLogRequest, opts ...grpc.CallOption) (*GetLogReply, error) {
	out := new(GetLogReply)
	err := c.cc.Invoke(ctx, Log_GetLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logClient) ListLog(ctx context.Context, in *ListLogRequest, opts ...grpc.CallOption) (*ListLogReply, error) {
	out := new(ListLogReply)
	err := c.cc.Invoke(ctx, Log_ListLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogServer is the server API for Log service.
// All implementations must embed UnimplementedLogServer
// for forward compatibility
type LogServer interface {
	// GetLog 获取指定的日志信息
	GetLog(context.Context, *GetLogRequest) (*GetLogReply, error)
	// ListLog 获取日志信息列表
	ListLog(context.Context, *ListLogRequest) (*ListLogReply, error)
	mustEmbedUnimplementedLogServer()
}

// UnimplementedLogServer must be embedded to have forward compatible implementations.
type UnimplementedLogServer struct {
}

func (UnimplementedLogServer) GetLog(context.Context, *GetLogRequest) (*GetLogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLog not implemented")
}
func (UnimplementedLogServer) ListLog(context.Context, *ListLogRequest) (*ListLogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLog not implemented")
}
func (UnimplementedLogServer) mustEmbedUnimplementedLogServer() {}

// UnsafeLogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogServer will
// result in compilation errors.
type UnsafeLogServer interface {
	mustEmbedUnimplementedLogServer()
}

func RegisterLogServer(s grpc.ServiceRegistrar, srv LogServer) {
	s.RegisterService(&Log_ServiceDesc, srv)
}

func _Log_GetLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServer).GetLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Log_GetLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServer).GetLog(ctx, req.(*GetLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Log_ListLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServer).ListLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Log_ListLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServer).ListLog(ctx, req.(*ListLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Log_ServiceDesc is the grpc.ServiceDesc for Log service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Log_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cron.api.server.cron.log.v1.Log",
	HandlerType: (*LogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLog",
			Handler:    _Log_GetLog_Handler,
		},
		{
			MethodName: "ListLog",
			Handler:    _Log_ListLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/cron/server/log/cron_log_service.proto",
}