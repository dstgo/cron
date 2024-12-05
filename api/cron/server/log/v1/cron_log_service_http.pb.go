// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v4.24.4
// source: api/cron/server/log/cron_log_service.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationLogGetLog = "/cron.api.server.cron.log.v1.Log/GetLog"
const OperationLogListLog = "/cron.api.server.cron.log.v1.Log/ListLog"

type LogHTTPServer interface {
	// GetLog GetLog 获取指定的日志信息
	GetLog(context.Context, *GetLogRequest) (*GetLogReply, error)
	// ListLog ListLog 获取日志信息列表
	ListLog(context.Context, *ListLogRequest) (*ListLogReply, error)
}

func RegisterLogHTTPServer(s *http.Server, srv LogHTTPServer) {
	r := s.Route("/")
	r.GET("/cron/api/v1/log", _Log_GetLog0_HTTP_Handler(srv))
	r.GET("/cron/api/v1/logs", _Log_ListLog0_HTTP_Handler(srv))
}

func _Log_GetLog0_HTTP_Handler(srv LogHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetLogRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLogGetLog)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.GetLog(ctx, req.(*GetLogRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetLogReply)
		return ctx.Result(200, reply)
	}
}

func _Log_ListLog0_HTTP_Handler(srv LogHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListLogRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLogListLog)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.ListLog(ctx, req.(*ListLogRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListLogReply)
		return ctx.Result(200, reply)
	}
}

type LogHTTPClient interface {
	GetLog(ctx context.Context, req *GetLogRequest, opts ...http.CallOption) (rsp *GetLogReply, err error)
	ListLog(ctx context.Context, req *ListLogRequest, opts ...http.CallOption) (rsp *ListLogReply, err error)
}

type LogHTTPClientImpl struct {
	cc *http.Client
}

func NewLogHTTPClient(client *http.Client) LogHTTPClient {
	return &LogHTTPClientImpl{client}
}

func (c *LogHTTPClientImpl) GetLog(ctx context.Context, in *GetLogRequest, opts ...http.CallOption) (*GetLogReply, error) {
	var out GetLogReply
	pattern := "/cron/api/v1/log"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationLogGetLog))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LogHTTPClientImpl) ListLog(ctx context.Context, in *ListLogRequest, opts ...http.CallOption) (*ListLogReply, error) {
	var out ListLogReply
	pattern := "/cron/api/v1/logs"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationLogListLog))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
