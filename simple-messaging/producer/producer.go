package producer

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Producer() {
	broker := "localhost:9092"
	topic := "kafka-test"

	p, err := kafka.NewProducer(
		&kafka.ConfigMap{
			"bootstrap.servers": broker,
		},
	)
	if err != nil {
		log.Fatalf("Failed to create producer: %s", err)
	}

	defer p.Close()
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition.Error)
				} else {
					fmt.Printf("Message delivered to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	message := "Hello from Kafka Online Test"
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic: &topic, Partition: kafka.PartitionAny,
		}, Value: []byte(message)}, nil)
	if err != nil {
		log.Fatalf("Failed to produce message: %s", err)
	}
}
