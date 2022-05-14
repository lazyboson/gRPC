package main

import (
	"context"
	"fmt"
	"log"

	pb "gRPC-test/gen/pb"

	"google.golang.org/grpc"
)

type Employee struct {
	Name string
	Salary uint32
	Address string
}

type Employees struct {
	Employs []Employee
}


func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	client := pb.NewEmployeeAPIClient(conn)
	
	resp, err := client.AddEmployee(context.Background(), &pb.Employee{Name: "amar", Salary: 50000, Address: "New Delhi India 110002" })
	if err != nil {
		log.Println(err)
	}
	fmt.Println(resp)
	fmt.Println(resp.Message)
	ep1 := &pb.Employee{Name: "ashutosh", Salary: 2000, Address: "Lucknow India"}
	ep2 := &pb.Employee{Name: "priya", Salary: 20000, Address: "UK India"}
	ep3 := &pb.Employee{Name: "rahul", Salary: 200000, Address: "Chennai India"}
	

	data := &pb.Employees {
		Employees: []*pb.Employee {ep1, ep2, ep3},
	}

	resp, err = client.PopulateEmployees(context.Background(), data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}