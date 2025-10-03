package kafka_producer

import (
	entity "auth-service/domain/entity"
	interfaces "auth-service/domain/interface"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProducer struct {
	Producer *kafka.Producer
}

type KafkaProducerConfig struct {
	Broker           string
	SecurityProtocol string
	SaslMechanism    string
	Username         string
	Password         string
}

func NewKafkaProducer(config KafkaProducerConfig) (interfaces.Producer, error) {
	protocol := config.SecurityProtocol
	saslMechanism := config.SaslMechanism
	if protocol == "" {
		protocol = "PLAINTEXT"
	}
	if saslMechanism == "" {
		saslMechanism = "PLAIN"
	}

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": config.Broker,
		"security.protocol": protocol,
		"sasl.mechanism":    saslMechanism,
		"sasl.username":     config.Username,
		"sasl.password":     config.Password,
	})
	if err != nil {
		return nil, err
	}

	kp := &KafkaProducer{Producer: p}

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Failed to deliver message: %v\n", ev.TopicPartition.Error)
				} else {
					log.Printf("Message delivered to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	return kp, nil
}

func (kp *KafkaProducer) Produce(msg entity.Message) error {
	return kp.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &msg.Topic, Partition: kafka.PartitionAny},
		Value:          []byte(msg.Content),
	}, nil)
}

func (kp *KafkaProducer) Close() {
	kp.Producer.Flush(15 * 1000)
	kp.Producer.Close()
}
