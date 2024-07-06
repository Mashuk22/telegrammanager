package controller

import (
	"log"

	"github.com/Mashuk22/telegrammanager/pkg/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Controller struct {
	UserServiceClient userpb.UserServiceClient
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) InitGRPCClients() {
	conn, err := grpc.NewClient("server:7077", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("NewClient faild: %v", err)
	}
	defer conn.Close()

	c.UserServiceClient = userpb.NewUserServiceClient(conn)
}

type Message struct {
	Message string `json:"message" example:"message"`
}
