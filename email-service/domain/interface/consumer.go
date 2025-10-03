package interfaces

type Consumer interface {
	Consume(func(topic string, message string)) error
	Close() error
}
