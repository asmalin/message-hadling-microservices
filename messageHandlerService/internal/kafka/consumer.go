package kafka

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"message-handler/internal/model"
	"message-handler/internal/service"

	"github.com/segmentio/kafka-go"
)

func Listen(service service.Message) {

	ctx := context.Background()

	l := log.New(os.Stdout, "kafka reader: ", 0)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{os.Getenv("KAFKA_1")},
		Topic:   os.Getenv("PRODUCER_TOPIC_NAME"),
		MaxWait: 3 * time.Second,
		Logger:  l,
	})

	for {

		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}

		go func() {
			var message model.Message
			err := json.Unmarshal(msg.Value, &message)
			if err != nil {
				log.Printf("Unmarshal JSON error: %v", err)
				return
			}

			err = service.ProcessMessage(message)
			if err != nil {
				log.Fatalf(err.Error())
			}
		}()

	}

}
