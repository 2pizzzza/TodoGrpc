// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.19.6
// source: proto/todo_v1/service.proto

package todo_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	TodoV1_CreateTodo_FullMethodName   = "/todo_v1.TodoV1/CreateTodo"
	TodoV1_CompleteTodo_FullMethodName = "/todo_v1.TodoV1/CompleteTodo"
	TodoV1_GetTodoById_FullMethodName  = "/todo_v1.TodoV1/GetTodoById"
	TodoV1_RemoveTodo_FullMethodName   = "/todo_v1.TodoV1/RemoveTodo"
	TodoV1_ChangeTodo_FullMethodName   = "/todo_v1.TodoV1/ChangeTodo"
	TodoV1_GetAllTodo_FullMethodName   = "/todo_v1.TodoV1/GetAllTodo"
)

// TodoV1Client is the client API for TodoV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Сервис для управления задачами
type TodoV1Client interface {
	// Создание новой задачи
	CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*TodoResponse, error)
	// Завершение задачи по идентификатору
	CompleteTodo(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*TodoResponse, error)
	// Получение задачи по идентификатору
	GetTodoById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*TodoResponse, error)
	// Удаление задачи по идентификатору
	RemoveTodo(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*InfoResponse, error)
	// Изменение существующей задачи
	ChangeTodo(ctx context.Context, in *ChangeTodoRequest, opts ...grpc.CallOption) (*TodoResponse, error)
	GetAllTodo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TodosResponse, error)
}

type todoV1Client struct {
	cc grpc.ClientConnInterface
}

func NewTodoV1Client(cc grpc.ClientConnInterface) TodoV1Client {
	return &todoV1Client{cc}
}

func (c *todoV1Client) CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*TodoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TodoResponse)
	err := c.cc.Invoke(ctx, TodoV1_CreateTodo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoV1Client) CompleteTodo(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*TodoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TodoResponse)
	err := c.cc.Invoke(ctx, TodoV1_CompleteTodo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoV1Client) GetTodoById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*TodoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TodoResponse)
	err := c.cc.Invoke(ctx, TodoV1_GetTodoById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoV1Client) RemoveTodo(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*InfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InfoResponse)
	err := c.cc.Invoke(ctx, TodoV1_RemoveTodo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoV1Client) ChangeTodo(ctx context.Context, in *ChangeTodoRequest, opts ...grpc.CallOption) (*TodoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TodoResponse)
	err := c.cc.Invoke(ctx, TodoV1_ChangeTodo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoV1Client) GetAllTodo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TodosResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TodosResponse)
	err := c.cc.Invoke(ctx, TodoV1_GetAllTodo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoV1Server is the server API for TodoV1 service.
// All implementations must embed UnimplementedTodoV1Server
// for forward compatibility
//
// Сервис для управления задачами
type TodoV1Server interface {
	// Создание новой задачи
	CreateTodo(context.Context, *CreateTodoRequest) (*TodoResponse, error)
	// Завершение задачи по идентификатору
	CompleteTodo(context.Context, *IdRequest) (*TodoResponse, error)
	// Получение задачи по идентификатору
	GetTodoById(context.Context, *IdRequest) (*TodoResponse, error)
	// Удаление задачи по идентификатору
	RemoveTodo(context.Context, *IdRequest) (*InfoResponse, error)
	// Изменение существующей задачи
	ChangeTodo(context.Context, *ChangeTodoRequest) (*TodoResponse, error)
	GetAllTodo(context.Context, *emptypb.Empty) (*TodosResponse, error)
	mustEmbedUnimplementedTodoV1Server()
}

// UnimplementedTodoV1Server must be embedded to have forward compatible implementations.
type UnimplementedTodoV1Server struct {
}

func (UnimplementedTodoV1Server) CreateTodo(context.Context, *CreateTodoRequest) (*TodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodo not implemented")
}
func (UnimplementedTodoV1Server) CompleteTodo(context.Context, *IdRequest) (*TodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteTodo not implemented")
}
func (UnimplementedTodoV1Server) GetTodoById(context.Context, *IdRequest) (*TodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodoById not implemented")
}
func (UnimplementedTodoV1Server) RemoveTodo(context.Context, *IdRequest) (*InfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveTodo not implemented")
}
func (UnimplementedTodoV1Server) ChangeTodo(context.Context, *ChangeTodoRequest) (*TodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeTodo not implemented")
}
func (UnimplementedTodoV1Server) GetAllTodo(context.Context, *emptypb.Empty) (*TodosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTodo not implemented")
}
func (UnimplementedTodoV1Server) mustEmbedUnimplementedTodoV1Server() {}

// UnsafeTodoV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoV1Server will
// result in compilation errors.
type UnsafeTodoV1Server interface {
	mustEmbedUnimplementedTodoV1Server()
}

func RegisterTodoV1Server(s grpc.ServiceRegistrar, srv TodoV1Server) {
	s.RegisterService(&TodoV1_ServiceDesc, srv)
}

func _TodoV1_CreateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoV1Server).CreateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoV1_CreateTodo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoV1Server).CreateTodo(ctx, req.(*CreateTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoV1_CompleteTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoV1Server).CompleteTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoV1_CompleteTodo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoV1Server).CompleteTodo(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoV1_GetTodoById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoV1Server).GetTodoById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoV1_GetTodoById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoV1Server).GetTodoById(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoV1_RemoveTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoV1Server).RemoveTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoV1_RemoveTodo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoV1Server).RemoveTodo(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoV1_ChangeTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoV1Server).ChangeTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoV1_ChangeTodo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoV1Server).ChangeTodo(ctx, req.(*ChangeTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoV1_GetAllTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoV1Server).GetAllTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoV1_GetAllTodo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoV1Server).GetAllTodo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// TodoV1_ServiceDesc is the grpc.ServiceDesc for TodoV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodoV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "todo_v1.TodoV1",
	HandlerType: (*TodoV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTodo",
			Handler:    _TodoV1_CreateTodo_Handler,
		},
		{
			MethodName: "CompleteTodo",
			Handler:    _TodoV1_CompleteTodo_Handler,
		},
		{
			MethodName: "GetTodoById",
			Handler:    _TodoV1_GetTodoById_Handler,
		},
		{
			MethodName: "RemoveTodo",
			Handler:    _TodoV1_RemoveTodo_Handler,
		},
		{
			MethodName: "ChangeTodo",
			Handler:    _TodoV1_ChangeTodo_Handler,
		},
		{
			MethodName: "GetAllTodo",
			Handler:    _TodoV1_GetAllTodo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/todo_v1/service.proto",
}
