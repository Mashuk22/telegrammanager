package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/IBM/sarama"
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

type LogEntry struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}

func main() {

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	db.New(conn)

	kafkaBrokers := os.Getenv("KAFKA_BROKERS")

	// Создание Kafka producer
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{kafkaBrokers}, config)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer producer.Close()

	// Отправка логов в Kafka
	level := "info"
	message := "This is a sample log entry."
	logEntry := LogEntry{
		Level:   level,
		Message: message,
	}

	data, err := json.Marshal(logEntry)
	if err != nil {
		log.Fatalf("Failed to marshal log entry: %v", err)
	}

	msg := &sarama.ProducerMessage{
		Topic: "log-topic",
		Value: sarama.ByteEncoder(data),
	}

	_, _, err = producer.SendMessage(msg)
	if err != nil {
		log.Fatalf("Failed to send log entry to Kafka: %v", err)
	}

	log.Printf("Sent log entry to Kafka: %s", message)

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
