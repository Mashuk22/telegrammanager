package main

import (
	"github.com/Mashuk22/telegrammanager/apigateway/controller"
	_ "github.com/Mashuk22/telegrammanager/apigateway/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	c := controller.NewController()
	c.InitGRPCClients()

	ginEngine := gin.Default()

	ginEngine.GET("/users", c.GetUsersHandler)
	ginEngine.GET("/users/:id", c.GetUserHandler)
	ginEngine.POST("/users", c.CreateUserHandler)

	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ginEngine.Run(":8080")
}
