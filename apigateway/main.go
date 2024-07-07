package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/Mashuk22/telegrammanager/apigateway/controller"
	_ "github.com/Mashuk22/telegrammanager/apigateway/docs"
	"github.com/gin-gonic/gin"
	swagfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	exp, err := otlptracehttp.New(ctx, otlptracehttp.WithEndpointURL("http://otel-collector:4318/v1/traces"))
	if err != nil {
		panic(err)
	}

	controller.TracerProvider = trace.NewTracerProvider(trace.WithBatcher(exp))
	defer func() {
		if err := controller.TracerProvider.Shutdown(ctx); err != nil {
			panic(err)
		}
	}()
	otel.SetTracerProvider(controller.TracerProvider)

	c := controller.NewController()
	c.InitGRPCClients()

	ginEngine := gin.Default()
	ginEngine.Use(otelgin.Middleware("apigateway"))

	ginEngine.GET("/users", c.GetUsersHandler)
	ginEngine.GET("/users/:id", c.GetUserHandler)
	ginEngine.POST("/users", c.CreateUserHandler)

	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swagfiles.Handler))

	srvErr := make(chan error, 1)
	go func() {
		srvErr <- ginEngine.Run(":7078")
	}()

	select {
	case <-srvErr:
		return err
	case <-ctx.Done():
		stop()
	}
	return err
}
