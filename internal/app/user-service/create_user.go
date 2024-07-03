package userservice

import (
	"context"

	userpb "github.com/Mashuk22/telegrammanager/pkg/user_service"
)

func CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {

	return &userpb.CreateUserResponse{}, nil
}
