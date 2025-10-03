package kafka_consumer

import (
	"context"
	interfaces "email-service/domain/interface"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	Consumer *kafka.Consumer
	Topics   []string
}

type KafkaConsumerConfig struct {
	Broker           string
	SecurityProtocol string
	SaslMechanism    string
	Username         string
	Password         string
	Topics           []string
	GroupID          string
}

func NewKafkaConsumer(config KafkaConsumerConfig) (interfaces.Consumer, error) {
	protocol := config.SecurityProtocol
	saslMechanism := config.SaslMechanism
	if protocol == "" {
		protocol = "PLAINTEXT"
	}
	if saslMechanism == "" {
		saslMechanism = "PLAIN"
	}

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.Broker,
		"group.id":          config.GroupID,
		"auto.offset.reset": "earliest",
		"security.protocol": protocol,
		"sasl.mechanism":    saslMechanism,
		"sasl.username":     config.Username,
		"sasl.password":     config.Password,
	})

	if err != nil {
		return nil, err
	}

	err = c.SubscribeTopics(config.Topics, nil)
	if err != nil {
		return nil, err
	}

	kc := &KafkaConsumer{Consumer: c, Topics: config.Topics}

	return kc, nil

}

func (kc *KafkaConsumer) Consume(handler func(topic string, message string)) error {

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
		<-sigchan
		cancel()
	}()

	fmt.Println("Consumer started. Waiting for messages...")

	run := true
	for run {
		select {

		case <-ctx.Done():
			fmt.Println("Shutting down consumer...")
			return nil
		default:
			msg, err := kc.Consumer.ReadMessage(100 * time.Millisecond)
			if err == nil {
				handler(*msg.TopicPartition.Topic, string(msg.Value))
			} else {
				if err.Error() != "Local: Timed out" {
					fmt.Printf("Consumer error: %v\n", err)
				}
			}
		}
	}

	return nil
}

func (kc *KafkaConsumer) Close() error {
	return kc.Consumer.Close()
}
