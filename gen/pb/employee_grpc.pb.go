// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: employee.proto

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

// EmployeeAPIClient is the client API for EmployeeAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmployeeAPIClient interface {
	AddEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*Response, error)
	PopulateEmployees(ctx context.Context, in *Employees, opts ...grpc.CallOption) (*Response, error)
}

type employeeAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewEmployeeAPIClient(cc grpc.ClientConnInterface) EmployeeAPIClient {
	return &employeeAPIClient{cc}
}

func (c *employeeAPIClient) AddEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/main.EmployeeAPI/AddEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeAPIClient) PopulateEmployees(ctx context.Context, in *Employees, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/main.EmployeeAPI/PopulateEmployees", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmployeeAPIServer is the server API for EmployeeAPI service.
// All implementations must embed UnimplementedEmployeeAPIServer
// for forward compatibility
type EmployeeAPIServer interface {
	AddEmployee(context.Context, *Employee) (*Response, error)
	PopulateEmployees(context.Context, *Employees) (*Response, error)
	mustEmbedUnimplementedEmployeeAPIServer()
}

// UnimplementedEmployeeAPIServer must be embedded to have forward compatible implementations.
type UnimplementedEmployeeAPIServer struct {
}

func (UnimplementedEmployeeAPIServer) AddEmployee(context.Context, *Employee) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEmployee not implemented")
}
func (UnimplementedEmployeeAPIServer) PopulateEmployees(context.Context, *Employees) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PopulateEmployees not implemented")
}
func (UnimplementedEmployeeAPIServer) mustEmbedUnimplementedEmployeeAPIServer() {}

// UnsafeEmployeeAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmployeeAPIServer will
// result in compilation errors.
type UnsafeEmployeeAPIServer interface {
	mustEmbedUnimplementedEmployeeAPIServer()
}

func RegisterEmployeeAPIServer(s grpc.ServiceRegistrar, srv EmployeeAPIServer) {
	s.RegisterService(&EmployeeAPI_ServiceDesc, srv)
}

func _EmployeeAPI_AddEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Employee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeAPIServer).AddEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.EmployeeAPI/AddEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeAPIServer).AddEmployee(ctx, req.(*Employee))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeAPI_PopulateEmployees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Employees)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeAPIServer).PopulateEmployees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.EmployeeAPI/PopulateEmployees",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeAPIServer).PopulateEmployees(ctx, req.(*Employees))
	}
	return interceptor(ctx, in, info, handler)
}

// EmployeeAPI_ServiceDesc is the grpc.ServiceDesc for EmployeeAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmployeeAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.EmployeeAPI",
	HandlerType: (*EmployeeAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddEmployee",
			Handler:    _EmployeeAPI_AddEmployee_Handler,
		},
		{
			MethodName: "PopulateEmployees",
			Handler:    _EmployeeAPI_PopulateEmployees_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "employee.proto",
}
