package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/Mashuk22/telegrammanager/userservice/db"
	"github.com/Mashuk22/telegrammanager/userservice/internal/app/userservice"
	"github.com/Mashuk22/telegrammanager/userservice/internal/rabbitmq"
	"github.com/Mashuk22/telegrammanager/userservice/pkg/userpb"
	"google.golang.org/grpc"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://postgres:mypassword@localhost:6543/telegrammanager?sslmode=disable"
)

func main() {

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	db.New(conn)

	server := grpc.NewServer()
	userpb.RegisterUserServiceServer(server, &userservice.Server{})
	listener, err := net.Listen("tcp", "0.0.0.0:7077")
	if err != nil {
		log.Fatal(err)
	}

	rmq := rabbitmq.NewService()
	err = rmq.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer rmq.Conn.Close()

	err = rmq.Publish("Hi")
	if err != nil {
		log.Fatal(err)
	}

	go rmq.Consume()

	fmt.Println("Starting gRPC server...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failer еo serve %v", err)
	}

}
