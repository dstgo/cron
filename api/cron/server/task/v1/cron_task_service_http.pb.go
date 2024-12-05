// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v4.24.4
// source: api/cron/server/task/cron_task_service.proto

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

const OperationTaskCancelExecTask = "/cron.api.server.cron.task.v1.Task/CancelExecTask"
const OperationTaskCreateTask = "/cron.api.server.cron.task.v1.Task/CreateTask"
const OperationTaskCreateTaskGroup = "/cron.api.server.cron.task.v1.Task/CreateTaskGroup"
const OperationTaskDeleteTask = "/cron.api.server.cron.task.v1.Task/DeleteTask"
const OperationTaskDeleteTaskGroup = "/cron.api.server.cron.task.v1.Task/DeleteTaskGroup"
const OperationTaskExecTask = "/cron.api.server.cron.task.v1.Task/ExecTask"
const OperationTaskGetTask = "/cron.api.server.cron.task.v1.Task/GetTask"
const OperationTaskGetTaskGroup = "/cron.api.server.cron.task.v1.Task/GetTaskGroup"
const OperationTaskListTask = "/cron.api.server.cron.task.v1.Task/ListTask"
const OperationTaskListTaskGroup = "/cron.api.server.cron.task.v1.Task/ListTaskGroup"
const OperationTaskUpdateTask = "/cron.api.server.cron.task.v1.Task/UpdateTask"
const OperationTaskUpdateTaskGroup = "/cron.api.server.cron.task.v1.Task/UpdateTaskGroup"
const OperationTaskUpdateTaskStatus = "/cron.api.server.cron.task.v1.Task/UpdateTaskStatus"

type TaskHTTPServer interface {
	// CancelExecTask CancelExecTask 取消指定任务
	CancelExecTask(context.Context, *CancelExecTaskRequest) (*CancelExecTaskReply, error)
	// CreateTask CreateTask 创建任务信息
	CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskReply, error)
	// CreateTaskGroup CreateTaskGroup 创建任务分组
	CreateTaskGroup(context.Context, *CreateTaskGroupRequest) (*CreateTaskGroupReply, error)
	// DeleteTask DeleteTask 删除任务信息
	DeleteTask(context.Context, *DeleteTaskRequest) (*DeleteTaskReply, error)
	// DeleteTaskGroup DeleteTaskGroup 删除任务分组
	DeleteTaskGroup(context.Context, *DeleteTaskGroupRequest) (*DeleteTaskGroupReply, error)
	// ExecTask ExecTask 立即执行任务
	ExecTask(context.Context, *ExecTaskRequest) (*ExecTaskReply, error)
	// GetTask GetTask 获取指定的任务信息
	GetTask(context.Context, *GetTaskRequest) (*GetTaskReply, error)
	// GetTaskGroup GetTaskGroup 获取指定的任务分组
	GetTaskGroup(context.Context, *GetTaskGroupRequest) (*GetTaskGroupReply, error)
	// ListTask ListTask 获取任务信息列表
	ListTask(context.Context, *ListTaskRequest) (*ListTaskReply, error)
	// ListTaskGroup ListTaskGroup 获取任务分组列表
	ListTaskGroup(context.Context, *ListTaskGroupRequest) (*ListTaskGroupReply, error)
	// UpdateTask UpdateTask 更新任务信息
	UpdateTask(context.Context, *UpdateTaskRequest) (*UpdateTaskReply, error)
	// UpdateTaskGroup UpdateTaskGroup 更新任务分组
	UpdateTaskGroup(context.Context, *UpdateTaskGroupRequest) (*UpdateTaskGroupReply, error)
	// UpdateTaskStatus UpdateTaskStatus 更新任务信息状态
	UpdateTaskStatus(context.Context, *UpdateTaskStatusRequest) (*UpdateTaskStatusReply, error)
}

func RegisterTaskHTTPServer(s *http.Server, srv TaskHTTPServer) {
	r := s.Route("/")
	r.GET("/cron/api/v1/task_group", _Task_GetTaskGroup0_HTTP_Handler(srv))
	r.GET("/cron/api/v1/task_groups", _Task_ListTaskGroup0_HTTP_Handler(srv))
	r.POST("/cron/api/v1/task_group", _Task_CreateTaskGroup0_HTTP_Handler(srv))
	r.PUT("/cron/api/v1/task_group", _Task_UpdateTaskGroup0_HTTP_Handler(srv))
	r.DELETE("/cron/api/v1/task_group", _Task_DeleteTaskGroup0_HTTP_Handler(srv))
	r.GET("/cron/api/v1/task", _Task_GetTask0_HTTP_Handler(srv))
	r.GET("/cron/api/v1/tasks", _Task_ListTask0_HTTP_Handler(srv))
	r.POST("/cron/api/v1/task", _Task_CreateTask0_HTTP_Handler(srv))
	r.PUT("/cron/api/v1/task", _Task_UpdateTask0_HTTP_Handler(srv))
	r.PUT("/cron/api/v1/task/status", _Task_UpdateTaskStatus0_HTTP_Handler(srv))
	r.DELETE("/cron/api/v1/task", _Task_DeleteTask0_HTTP_Handler(srv))
	r.POST("/cron/api/v1/task/exec", _Task_ExecTask0_HTTP_Handler(srv))
	r.POST("/cron/api/v1/task/cancel", _Task_CancelExecTask0_HTTP_Handler(srv))
}

func _Task_GetTaskGroup0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTaskGroupRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskGetTaskGroup)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.GetTaskGroup(ctx, req.(*GetTaskGroupRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetTaskGroupReply)
		return ctx.Result(200, reply)
	}
}

func _Task_ListTaskGroup0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListTaskGroupRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskListTaskGroup)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.ListTaskGroup(ctx, req.(*ListTaskGroupRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListTaskGroupReply)
		return ctx.Result(200, reply)
	}
}

func _Task_CreateTaskGroup0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateTaskGroupRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskCreateTaskGroup)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.CreateTaskGroup(ctx, req.(*CreateTaskGroupRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateTaskGroupReply)
		return ctx.Result(200, reply)
	}
}

func _Task_UpdateTaskGroup0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateTaskGroupRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskUpdateTaskGroup)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.UpdateTaskGroup(ctx, req.(*UpdateTaskGroupRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateTaskGroupReply)
		return ctx.Result(200, reply)
	}
}

func _Task_DeleteTaskGroup0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteTaskGroupRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskDeleteTaskGroup)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.DeleteTaskGroup(ctx, req.(*DeleteTaskGroupRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteTaskGroupReply)
		return ctx.Result(200, reply)
	}
}

func _Task_GetTask0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTaskRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskGetTask)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.GetTask(ctx, req.(*GetTaskRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetTaskReply)
		return ctx.Result(200, reply)
	}
}

func _Task_ListTask0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListTaskRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskListTask)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.ListTask(ctx, req.(*ListTaskRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListTaskReply)
		return ctx.Result(200, reply)
	}
}

func _Task_CreateTask0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateTaskRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskCreateTask)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.CreateTask(ctx, req.(*CreateTaskRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateTaskReply)
		return ctx.Result(200, reply)
	}
}

func _Task_UpdateTask0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateTaskRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskUpdateTask)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.UpdateTask(ctx, req.(*UpdateTaskRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateTaskReply)
		return ctx.Result(200, reply)
	}
}

func _Task_UpdateTaskStatus0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateTaskStatusRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskUpdateTaskStatus)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.UpdateTaskStatus(ctx, req.(*UpdateTaskStatusRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateTaskStatusReply)
		return ctx.Result(200, reply)
	}
}

func _Task_DeleteTask0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteTaskRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskDeleteTask)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.DeleteTask(ctx, req.(*DeleteTaskRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteTaskReply)
		return ctx.Result(200, reply)
	}
}

func _Task_ExecTask0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ExecTaskRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskExecTask)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.ExecTask(ctx, req.(*ExecTaskRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ExecTaskReply)
		return ctx.Result(200, reply)
	}
}

func _Task_CancelExecTask0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CancelExecTaskRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskCancelExecTask)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.CancelExecTask(ctx, req.(*CancelExecTaskRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CancelExecTaskReply)
		return ctx.Result(200, reply)
	}
}

type TaskHTTPClient interface {
	CancelExecTask(ctx context.Context, req *CancelExecTaskRequest, opts ...http.CallOption) (rsp *CancelExecTaskReply, err error)
	CreateTask(ctx context.Context, req *CreateTaskRequest, opts ...http.CallOption) (rsp *CreateTaskReply, err error)
	CreateTaskGroup(ctx context.Context, req *CreateTaskGroupRequest, opts ...http.CallOption) (rsp *CreateTaskGroupReply, err error)
	DeleteTask(ctx context.Context, req *DeleteTaskRequest, opts ...http.CallOption) (rsp *DeleteTaskReply, err error)
	DeleteTaskGroup(ctx context.Context, req *DeleteTaskGroupRequest, opts ...http.CallOption) (rsp *DeleteTaskGroupReply, err error)
	ExecTask(ctx context.Context, req *ExecTaskRequest, opts ...http.CallOption) (rsp *ExecTaskReply, err error)
	GetTask(ctx context.Context, req *GetTaskRequest, opts ...http.CallOption) (rsp *GetTaskReply, err error)
	GetTaskGroup(ctx context.Context, req *GetTaskGroupRequest, opts ...http.CallOption) (rsp *GetTaskGroupReply, err error)
	ListTask(ctx context.Context, req *ListTaskRequest, opts ...http.CallOption) (rsp *ListTaskReply, err error)
	ListTaskGroup(ctx context.Context, req *ListTaskGroupRequest, opts ...http.CallOption) (rsp *ListTaskGroupReply, err error)
	UpdateTask(ctx context.Context, req *UpdateTaskRequest, opts ...http.CallOption) (rsp *UpdateTaskReply, err error)
	UpdateTaskGroup(ctx context.Context, req *UpdateTaskGroupRequest, opts ...http.CallOption) (rsp *UpdateTaskGroupReply, err error)
	UpdateTaskStatus(ctx context.Context, req *UpdateTaskStatusRequest, opts ...http.CallOption) (rsp *UpdateTaskStatusReply, err error)
}

type TaskHTTPClientImpl struct {
	cc *http.Client
}

func NewTaskHTTPClient(client *http.Client) TaskHTTPClient {
	return &TaskHTTPClientImpl{client}
}

func (c *TaskHTTPClientImpl) CancelExecTask(ctx context.Context, in *CancelExecTaskRequest, opts ...http.CallOption) (*CancelExecTaskReply, error) {
	var out CancelExecTaskReply
	pattern := "/cron/api/v1/task/cancel"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTaskCancelExecTask))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...http.CallOption) (*CreateTaskReply, error) {
	var out CreateTaskReply
	pattern := "/cron/api/v1/task"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTaskCreateTask))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) CreateTaskGroup(ctx context.Context, in *CreateTaskGroupRequest, opts ...http.CallOption) (*CreateTaskGroupReply, error) {
	var out CreateTaskGroupReply
	pattern := "/cron/api/v1/task_group"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTaskCreateTaskGroup))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...http.CallOption) (*DeleteTaskReply, error) {
	var out DeleteTaskReply
	pattern := "/cron/api/v1/task"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTaskDeleteTask))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) DeleteTaskGroup(ctx context.Context, in *DeleteTaskGroupRequest, opts ...http.CallOption) (*DeleteTaskGroupReply, error) {
	var out DeleteTaskGroupReply
	pattern := "/cron/api/v1/task_group"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTaskDeleteTaskGroup))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) ExecTask(ctx context.Context, in *ExecTaskRequest, opts ...http.CallOption) (*ExecTaskReply, error) {
	var out ExecTaskReply
	pattern := "/cron/api/v1/task/exec"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTaskExecTask))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) GetTask(ctx context.Context, in *GetTaskRequest, opts ...http.CallOption) (*GetTaskReply, error) {
	var out GetTaskReply
	pattern := "/cron/api/v1/task"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTaskGetTask))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) GetTaskGroup(ctx context.Context, in *GetTaskGroupRequest, opts ...http.CallOption) (*GetTaskGroupReply, error) {
	var out GetTaskGroupReply
	pattern := "/cron/api/v1/task_group"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTaskGetTaskGroup))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) ListTask(ctx context.Context, in *ListTaskRequest, opts ...http.CallOption) (*ListTaskReply, error) {
	var out ListTaskReply
	pattern := "/cron/api/v1/tasks"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTaskListTask))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) ListTaskGroup(ctx context.Context, in *ListTaskGroupRequest, opts ...http.CallOption) (*ListTaskGroupReply, error) {
	var out ListTaskGroupReply
	pattern := "/cron/api/v1/task_groups"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTaskListTaskGroup))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...http.CallOption) (*UpdateTaskReply, error) {
	var out UpdateTaskReply
	pattern := "/cron/api/v1/task"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTaskUpdateTask))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) UpdateTaskGroup(ctx context.Context, in *UpdateTaskGroupRequest, opts ...http.CallOption) (*UpdateTaskGroupReply, error) {
	var out UpdateTaskGroupReply
	pattern := "/cron/api/v1/task_group"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTaskUpdateTaskGroup))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) UpdateTaskStatus(ctx context.Context, in *UpdateTaskStatusRequest, opts ...http.CallOption) (*UpdateTaskStatusReply, error) {
	var out UpdateTaskStatusReply
	pattern := "/cron/api/v1/task/status"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTaskUpdateTaskStatus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
