package userservice

import (
	"context"

	"github.com/Mashuk22/telegrammanager/pkg/userpb"
)

type Server struct {
	userpb.UnimplementedUserServiceServer
}

func (server *Server) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {

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