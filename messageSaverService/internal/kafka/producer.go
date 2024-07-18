package kafka

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

func SendMessage(data interface{}) {
	brokers := []string{os.Getenv("KAFKA_1")}

	l := log.New(os.Stdout, "kafka writer: ", 0)

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  brokers,
		Topic:    os.Getenv("PRODUCER_TOPIC_NAME"),
		Balancer: &kafka.LeastBytes{},
		Logger:   l,
	})

	defer writer.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to encode to JSON: %v", err)
	}

	message := kafka.Message{
		Value: jsonData,
	}

	if writer.WriteMessages(context.Background(), message) != nil {
		log.Fatalf("could not write message %v", err)
	}

	writer.Close()
}
