package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Mashuk22/telegrammanager/pkg/userpb"
	"github.com/gin-gonic/gin"
)

// GetUsersHandler godoc
//
//	@Summary		Get Users
//	@Description	Get Users
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.User
//	@Router			/users [get]
func (c *Controller) GetUsersHandler(ctx *gin.Context) {
	resp, err := c.UserServiceClient.ListUsers(context.Background(), &userpb.ListUsersRequest{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *Controller) GetUserHandler(ctx *gin.Context) {

	fmt.Printf("getUserHandler response")
}

// CreateUserHandler godoc
//
//	@Summary		CreateUserHandler
//	@Description	CreateUserHandler
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		model.CreateUser		true	"Account ID"
//	@Success		200	{object}	model.User
//	@Router			/users [post]
func (c *Controller) CreateUserHandler(ctx *gin.Context) {
	var req userpb.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := c.UserServiceClient.CreateUser(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}
