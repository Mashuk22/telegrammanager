package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Mashuk22/telegrammanager/pkg/userpb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	userServiceClient userpb.UserServiceClient
)

func main() {
	initGRPCClients()

	ginEngine := gin.Default()

	ginEngine.GET("/users", getUsersHandler)
	ginEngine.GET("/users/:id", getUserHandler)
	ginEngine.POST("/users", createUserHandler)

	ginEngine.Run()
}

func getUsersHandler(c *gin.Context) {
	fmt.Printf("getUsersHandler response")
}

func getUserHandler(c *gin.Context) {
	resp, err := userServiceClient.ListUsers(context.Background(), &userpb.ListUsersRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func createUserHandler(c *gin.Context) {
	var req userpb.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := userServiceClient.CreateUser(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func initGRPCClients() {
	conn, err := grpc.NewClient("server:7077", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("NewClient faild: %v", err)
	}
	defer conn.Close()

	userServiceClient = userpb.NewUserServiceClient(conn)
}
