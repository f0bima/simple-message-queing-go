package consumer

import (
	"fmt"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
)

func Consumer() {
	fmt.Println("Running kafka consumer")
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	broker := os.Getenv("BROKER_URL")
	groupID := os.Getenv("CONSUMER_ID")

	topics := []string{"forgot-password", "otp"}

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		fmt.Printf("Failed to create consumer: %s", err)
	}
	defer c.Close()

	err = c.SubscribeTopics(topics, nil)
	if err != nil {
		fmt.Printf("Failed to subscribe to topic %s: %s", topics, err)
	}

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			topic := msg.TopicPartition.Topic

			switch *topic {
			case "otp":
				fmt.Printf("Sending otp to : %s\n", msg.Value)
			case "forgot-password":
				fmt.Printf("Sending reset password link to : %s\n", msg.Value)
			default:
				fmt.Printf("Unknown topic: %s\n", *topic)
			}

		} else {
			fmt.Printf("Consumer error: %v\n", err)
		}
	}

}
