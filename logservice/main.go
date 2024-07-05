package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
	"github.com/elastic/go-elasticsearch/v7"
)

type LogEntry struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}

func main() {
	log.Print("Start service")
	esHosts := os.Getenv("ELASTICSEARCH_HOSTS")
	kafkaBrokers := os.Getenv("KAFKA_BROKERS")

	// Подключение к Elasticsearch
	log.Print("Connection to Elasticsearch")
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{esHosts},
	})
	if err != nil {
		log.Fatalf("Failed to create Elasticsearch client: %v", err)
	}

	// Подключение к Kafka
	log.Print("Connection to Kafka")
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	consumer, err := sarama.NewConsumer([]string{kafkaBrokers}, config)
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}
	defer consumer.Close()

	// Обработка сообщений из Kafka
	log.Print("Kafka massage managing")
	partitionConsumer, err := consumer.ConsumePartition("log-topic", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to create partition consumer: %v", err)
	}
	defer partitionConsumer.Close()

	go func() {
		log.Print("Gorutines started")
		for msg := range partitionConsumer.Messages() {
			log.Printf("Get messge from kafka: %v", msg.Value)
			var logEntry LogEntry
			if err := json.Unmarshal(msg.Value, &logEntry); err != nil {
				log.Printf("Failed to unmarshal log entry: %v", err)
				continue
			}

			err := indexLogEntry(es, logEntry)
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	log.Println("Received interrupt signal, exiting...")
}

func indexLogEntry(es *elasticsearch.Client, logEntry LogEntry) error {
	data, err := json.Marshal(logEntry)
	if err != nil {
		return err
	}

	res, err := es.Index(
		"logs",
		bytes.NewReader(data),
		es.Index.WithContext(context.Background()),
		es.Index.WithDocumentType("_doc"),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}
