package userservice

import (
	"context"
	"log"

	"github.com/Mashuk22/telegrammanager/userservice/pkg/userpb"
)

type Server struct {
	userpb.UnimplementedUserServiceServer
}

func (server *Server) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	log.Print("CreateUser RPC call")
	return &userpb.CreateUserResponse{
		Id:           1,
		ChatId:       req.ChatId,
		Username:     req.Username,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		RoleId:       req.RoleId,
		IsSubscribed: req.IsSubscribed,
		CreatedAt:    "2023-07-04T12:00:00Z",
		UpdatedAt:    "2023-07-04T12:00:00Z",
	}, nil
}
