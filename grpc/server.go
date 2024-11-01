package main

import (
	"context"
	"errors"
	"fmt"
	"net"

	student_service "gituhb.com/ohno104dev/prac-go-grpc/grpc/idl/myproto"
	"google.golang.org/grpc"
)

type StudentServer struct {
	student_service.UnimplementedStudentServiceServer
}

func (StudentServer) GetStudentInfo(ctx context.Context, in *student_service.Request) (*student_service.Student, error) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("執行街口時出錯: %v\n", err)
		}
	}()
	if len(in.StudentId) == 0 {
		return nil, errors.New("學生ID不能為空")
	}

	student := &student_service.Student{
		Name:   "小白",
		Age:    3,
		Height: 1.52,
	}

	return student, nil
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:5678")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	student_service.RegisterStudentServiceServer(srv, new(StudentServer))

	err = srv.Serve(listen)
	if err != nil {
		panic(err)
	}

}
