// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.2
// source: isokitteh.proto

package pb

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

// RunnerManagerClient is the client API for RunnerManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RunnerManagerClient interface {
	Start(ctx context.Context, in *StartRunnerRequest, opts ...grpc.CallOption) (*StartRunnerResponse, error)
	RunnerHealth(ctx context.Context, in *RunnerHealthRequest, opts ...grpc.CallOption) (*RunnerHealthResponse, error)
	Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error)
	Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
}

type runnerManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewRunnerManagerClient(cc grpc.ClientConnInterface) RunnerManagerClient {
	return &runnerManagerClient{cc}
}

func (c *runnerManagerClient) Start(ctx context.Context, in *StartRunnerRequest, opts ...grpc.CallOption) (*StartRunnerResponse, error) {
	out := new(StartRunnerResponse)
	err := c.cc.Invoke(ctx, "/RunnerManager/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerManagerClient) RunnerHealth(ctx context.Context, in *RunnerHealthRequest, opts ...grpc.CallOption) (*RunnerHealthResponse, error) {
	out := new(RunnerHealthResponse)
	err := c.cc.Invoke(ctx, "/RunnerManager/RunnerHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerManagerClient) Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error) {
	out := new(StopResponse)
	err := c.cc.Invoke(ctx, "/RunnerManager/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerManagerClient) Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/RunnerManager/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RunnerManagerServer is the server API for RunnerManager service.
// All implementations must embed UnimplementedRunnerManagerServer
// for forward compatibility
type RunnerManagerServer interface {
	Start(context.Context, *StartRunnerRequest) (*StartRunnerResponse, error)
	RunnerHealth(context.Context, *RunnerHealthRequest) (*RunnerHealthResponse, error)
	Stop(context.Context, *StopRequest) (*StopResponse, error)
	Health(context.Context, *HealthRequest) (*HealthResponse, error)
	mustEmbedUnimplementedRunnerManagerServer()
}

// UnimplementedRunnerManagerServer must be embedded to have forward compatible implementations.
type UnimplementedRunnerManagerServer struct {
}

func (UnimplementedRunnerManagerServer) Start(context.Context, *StartRunnerRequest) (*StartRunnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedRunnerManagerServer) RunnerHealth(context.Context, *RunnerHealthRequest) (*RunnerHealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunnerHealth not implemented")
}
func (UnimplementedRunnerManagerServer) Stop(context.Context, *StopRequest) (*StopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedRunnerManagerServer) Health(context.Context, *HealthRequest) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedRunnerManagerServer) mustEmbedUnimplementedRunnerManagerServer() {}

// UnsafeRunnerManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RunnerManagerServer will
// result in compilation errors.
type UnsafeRunnerManagerServer interface {
	mustEmbedUnimplementedRunnerManagerServer()
}

func RegisterRunnerManagerServer(s grpc.ServiceRegistrar, srv RunnerManagerServer) {
	s.RegisterService(&RunnerManager_ServiceDesc, srv)
}

func _RunnerManager_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRunnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerManagerServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RunnerManager/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerManagerServer).Start(ctx, req.(*StartRunnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RunnerManager_RunnerHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunnerHealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerManagerServer).RunnerHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RunnerManager/RunnerHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerManagerServer).RunnerHealth(ctx, req.(*RunnerHealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RunnerManager_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerManagerServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RunnerManager/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerManagerServer).Stop(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RunnerManager_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerManagerServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RunnerManager/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerManagerServer).Health(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RunnerManager_ServiceDesc is the grpc.ServiceDesc for RunnerManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RunnerManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RunnerManager",
	HandlerType: (*RunnerManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _RunnerManager_Start_Handler,
		},
		{
			MethodName: "RunnerHealth",
			Handler:    _RunnerManager_RunnerHealth_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _RunnerManager_Stop_Handler,
		},
		{
			MethodName: "Health",
			Handler:    _RunnerManager_Health_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "isokitteh.proto",
}

// RunnerClient is the client API for Runner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RunnerClient interface {
	// Called at start of session
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
	// Execute a function in the runtime (skipped if it's a reply)
	Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error)
	// Reply from activity
	ActivityReply(ctx context.Context, in *ActivityReplyRequest, opts ...grpc.CallOption) (*ActivityReplyResponse, error)
	Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
}

type runnerClient struct {
	cc grpc.ClientConnInterface
}

func NewRunnerClient(cc grpc.ClientConnInterface) RunnerClient {
	return &runnerClient{cc}
}

func (c *runnerClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := c.cc.Invoke(ctx, "/Runner/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerClient) Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error) {
	out := new(ExecuteResponse)
	err := c.cc.Invoke(ctx, "/Runner/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerClient) ActivityReply(ctx context.Context, in *ActivityReplyRequest, opts ...grpc.CallOption) (*ActivityReplyResponse, error) {
	out := new(ActivityReplyResponse)
	err := c.cc.Invoke(ctx, "/Runner/ActivityReply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerClient) Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/Runner/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RunnerServer is the server API for Runner service.
// All implementations must embed UnimplementedRunnerServer
// for forward compatibility
type RunnerServer interface {
	// Called at start of session
	Start(context.Context, *StartRequest) (*StartResponse, error)
	// Execute a function in the runtime (skipped if it's a reply)
	Execute(context.Context, *ExecuteRequest) (*ExecuteResponse, error)
	// Reply from activity
	ActivityReply(context.Context, *ActivityReplyRequest) (*ActivityReplyResponse, error)
	Health(context.Context, *HealthRequest) (*HealthResponse, error)
	mustEmbedUnimplementedRunnerServer()
}

// UnimplementedRunnerServer must be embedded to have forward compatible implementations.
type UnimplementedRunnerServer struct {
}

func (UnimplementedRunnerServer) Start(context.Context, *StartRequest) (*StartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedRunnerServer) Execute(context.Context, *ExecuteRequest) (*ExecuteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedRunnerServer) ActivityReply(context.Context, *ActivityReplyRequest) (*ActivityReplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivityReply not implemented")
}
func (UnimplementedRunnerServer) Health(context.Context, *HealthRequest) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedRunnerServer) mustEmbedUnimplementedRunnerServer() {}

// UnsafeRunnerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RunnerServer will
// result in compilation errors.
type UnsafeRunnerServer interface {
	mustEmbedUnimplementedRunnerServer()
}

func RegisterRunnerServer(s grpc.ServiceRegistrar, srv RunnerServer) {
	s.RegisterService(&Runner_ServiceDesc, srv)
}

func _Runner_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Runner/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runner_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Runner/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).Execute(ctx, req.(*ExecuteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runner_ActivityReply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityReplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).ActivityReply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Runner/ActivityReply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).ActivityReply(ctx, req.(*ActivityReplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runner_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Runner/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).Health(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Runner_ServiceDesc is the grpc.ServiceDesc for Runner service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Runner_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Runner",
	HandlerType: (*RunnerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _Runner_Start_Handler,
		},
		{
			MethodName: "Execute",
			Handler:    _Runner_Execute_Handler,
		},
		{
			MethodName: "ActivityReply",
			Handler:    _Runner_ActivityReply_Handler,
		},
		{
			MethodName: "Health",
			Handler:    _Runner_Health_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "isokitteh.proto",
}

// WorkerClient is the client API for Worker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkerClient interface {
	// Runner starting activity
	Activity(ctx context.Context, in *ActivityRequest, opts ...grpc.CallOption) (*ActivityResponse, error)
	// Runner done with activity
	Done(ctx context.Context, in *DoneRequest, opts ...grpc.CallOption) (*DoneResponse, error)
	// Session logs
	Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error)
	// ak functions
	Sleep(ctx context.Context, in *SleepRequest, opts ...grpc.CallOption) (*SleepResponse, error)
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscribeResponse, error)
	NextEvent(ctx context.Context, in *NextEventRequest, opts ...grpc.CallOption) (*NextEventResponse, error)
	Unsubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*UnsubscribeResponse, error)
	Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
}

type workerClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkerClient(cc grpc.ClientConnInterface) WorkerClient {
	return &workerClient{cc}
}

func (c *workerClient) Activity(ctx context.Context, in *ActivityRequest, opts ...grpc.CallOption) (*ActivityResponse, error) {
	out := new(ActivityResponse)
	err := c.cc.Invoke(ctx, "/Worker/Activity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) Done(ctx context.Context, in *DoneRequest, opts ...grpc.CallOption) (*DoneResponse, error) {
	out := new(DoneResponse)
	err := c.cc.Invoke(ctx, "/Worker/Done", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error) {
	out := new(LogResponse)
	err := c.cc.Invoke(ctx, "/Worker/Log", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) Sleep(ctx context.Context, in *SleepRequest, opts ...grpc.CallOption) (*SleepResponse, error) {
	out := new(SleepResponse)
	err := c.cc.Invoke(ctx, "/Worker/Sleep", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscribeResponse, error) {
	out := new(SubscribeResponse)
	err := c.cc.Invoke(ctx, "/Worker/Subscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) NextEvent(ctx context.Context, in *NextEventRequest, opts ...grpc.CallOption) (*NextEventResponse, error) {
	out := new(NextEventResponse)
	err := c.cc.Invoke(ctx, "/Worker/NextEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) Unsubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*UnsubscribeResponse, error) {
	out := new(UnsubscribeResponse)
	err := c.cc.Invoke(ctx, "/Worker/Unsubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/Worker/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerServer is the server API for Worker service.
// All implementations must embed UnimplementedWorkerServer
// for forward compatibility
type WorkerServer interface {
	// Runner starting activity
	Activity(context.Context, *ActivityRequest) (*ActivityResponse, error)
	// Runner done with activity
	Done(context.Context, *DoneRequest) (*DoneResponse, error)
	// Session logs
	Log(context.Context, *LogRequest) (*LogResponse, error)
	// ak functions
	Sleep(context.Context, *SleepRequest) (*SleepResponse, error)
	Subscribe(context.Context, *SubscribeRequest) (*SubscribeResponse, error)
	NextEvent(context.Context, *NextEventRequest) (*NextEventResponse, error)
	Unsubscribe(context.Context, *UnsubscribeRequest) (*UnsubscribeResponse, error)
	Health(context.Context, *HealthRequest) (*HealthResponse, error)
	mustEmbedUnimplementedWorkerServer()
}

// UnimplementedWorkerServer must be embedded to have forward compatible implementations.
type UnimplementedWorkerServer struct {
}

func (UnimplementedWorkerServer) Activity(context.Context, *ActivityRequest) (*ActivityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Activity not implemented")
}
func (UnimplementedWorkerServer) Done(context.Context, *DoneRequest) (*DoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Done not implemented")
}
func (UnimplementedWorkerServer) Log(context.Context, *LogRequest) (*LogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Log not implemented")
}
func (UnimplementedWorkerServer) Sleep(context.Context, *SleepRequest) (*SleepResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sleep not implemented")
}
func (UnimplementedWorkerServer) Subscribe(context.Context, *SubscribeRequest) (*SubscribeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedWorkerServer) NextEvent(context.Context, *NextEventRequest) (*NextEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NextEvent not implemented")
}
func (UnimplementedWorkerServer) Unsubscribe(context.Context, *UnsubscribeRequest) (*UnsubscribeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unsubscribe not implemented")
}
func (UnimplementedWorkerServer) Health(context.Context, *HealthRequest) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedWorkerServer) mustEmbedUnimplementedWorkerServer() {}

// UnsafeWorkerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkerServer will
// result in compilation errors.
type UnsafeWorkerServer interface {
	mustEmbedUnimplementedWorkerServer()
}

func RegisterWorkerServer(s grpc.ServiceRegistrar, srv WorkerServer) {
	s.RegisterService(&Worker_ServiceDesc, srv)
}

func _Worker_Activity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Activity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Worker/Activity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Activity(ctx, req.(*ActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_Done_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Done(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Worker/Done",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Done(ctx, req.(*DoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_Log_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Log(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Worker/Log",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Log(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_Sleep_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SleepRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Sleep(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Worker/Sleep",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Sleep(ctx, req.(*SleepRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Worker/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Subscribe(ctx, req.(*SubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_NextEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NextEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).NextEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Worker/NextEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).NextEvent(ctx, req.(*NextEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_Unsubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Unsubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Worker/Unsubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Unsubscribe(ctx, req.(*UnsubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Worker/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Health(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Worker_ServiceDesc is the grpc.ServiceDesc for Worker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Worker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Worker",
	HandlerType: (*WorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Activity",
			Handler:    _Worker_Activity_Handler,
		},
		{
			MethodName: "Done",
			Handler:    _Worker_Done_Handler,
		},
		{
			MethodName: "Log",
			Handler:    _Worker_Log_Handler,
		},
		{
			MethodName: "Sleep",
			Handler:    _Worker_Sleep_Handler,
		},
		{
			MethodName: "Subscribe",
			Handler:    _Worker_Subscribe_Handler,
		},
		{
			MethodName: "NextEvent",
			Handler:    _Worker_NextEvent_Handler,
		},
		{
			MethodName: "Unsubscribe",
			Handler:    _Worker_Unsubscribe_Handler,
		},
		{
			MethodName: "Health",
			Handler:    _Worker_Health_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "isokitteh.proto",
}
