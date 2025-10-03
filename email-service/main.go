package main

import (
	"email-service/domain/entity"
	kafka_consumer "email-service/infrastructure/kafka"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Running kafka consumer")
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	topics := []string{"forgot-password", "otp"}

	config := kafka_consumer.KafkaConsumerConfig{
		Broker:           os.Getenv("BROKER_URL"),
		SecurityProtocol: os.Getenv("SECURITY_PROTOCOL"),
		SaslMechanism:    os.Getenv("SASL_MECHANISM"),
		Username:         os.Getenv("SASL_USERNAME"),
		Password:         os.Getenv("SASL_PASSWORD"),
		GroupID:          os.Getenv("CONSUMER_ID"),
		Topics:           topics,
	}

	consumer, err := kafka_consumer.NewKafkaConsumer(config)

	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %s\n", err)
		return
	}
	defer consumer.Close()

	err = consumer.Consume(func(topic string, message string) {
		switch topic {
		case "otp":
			var otp entity.Otp
			err := json.Unmarshal([]byte(message), &otp)
			if err != nil {
				log.Println("Failed to decode JSON:", err)
			} else {

				fmt.Printf("Sending otp %d to %s \n", otp.Otp, otp.Email)
			}
		case "forgot-password":
			fmt.Printf("Sending reset password link to : %s\n", message)
		default:
			fmt.Printf("Unknown topic: %s\n", topic)
		}
	})

	if err != nil {
		fmt.Printf("Failed to subscribe to topic %s: %s", topics, err)
	}

}
