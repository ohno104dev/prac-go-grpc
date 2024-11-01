package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	student_service "gituhb.com/ohno104dev/prac-go-grpc/grpc/idl/myproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestSomething(t *testing.T) {
	assert.True(t, true, "True is true!")
}

func TestGrpc(t *testing.T) {
	conn, err := grpc.NewClient(
		"127.0.0.1:5678",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		t.Fatalf("連接出錯: %v", err)
	}
	defer conn.Close()

	client := student_service.NewStudentServiceClient(conn)
	stu, err := client.GetStudentInfo(context.Background(), &student_service.Request{StudentId: "小白"})

	if err != nil {
		t.Fatalf("RPC 出錯: %v", err)
	}
	assert.NoError(t, err)
	fmt.Printf("1. 結果: %+v\n", stu)

	stu, err = client.GetStudentInfo(context.Background(), &student_service.Request{StudentId: ""})
	assert.Error(t, err)
	if err != nil {
		fmt.Printf("2. 空值測試 RPC 出錯: %v\n", err)
	}
}
