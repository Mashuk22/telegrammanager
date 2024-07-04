package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Mashuk22/telegrammanager/pkg/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("server:7077", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("NewClient faild: %v", err)
	}
	defer conn.Close()

	client := userpb.NewUserServiceClient(conn)
	resp, err := client.CreateUser(context.Background(), &userpb.CreateUserRequest{
		ChatId:       123,
		Username:     "testuser",
		FirstName:    "Test",
		LastName:     "User",
		RoleId:       1,
		IsSubscribed: true,
	})
	if err != nil {
		log.Fatalf("CreateUser failed: %v", err)
	}
	fmt.Printf("CreateUser response: %+v\n", resp)
}
