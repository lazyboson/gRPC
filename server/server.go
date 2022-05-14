package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "gRPC-test/gen/pb"

	"google.golang.org/grpc"
)

type EmployeeAPIServer struct {
	pb.UnimplementedEmployeeAPIServer
}

func (e *EmployeeAPIServer) AddEmployee(ctx context.Context, req *pb.Employee) (*pb.Response, error) {
	fmt.Println(req.Name)
	fmt.Println(req.Salary)
	fmt.Println(req.Address)
	return &pb.Response{Message: "Employee Added to DB", Error: "None"}, nil
}

func (e *EmployeeAPIServer) PopulateEmployees (ctx context.Context, req *pb.Employees) (*pb.Response, error) {
	employees := req.Employees
	for _, employee := range employees {
		fmt.Println(employee)
	}
	return &pb.Response{Message: "Success", Error: "None"}, nil
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterEmployeeAPIServer(grpcServer, &EmployeeAPIServer{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Println(err)
	}
}