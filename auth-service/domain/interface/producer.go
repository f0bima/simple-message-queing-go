package interfaces

import (
	entity "auth-service/domain/entity"
)

type Producer interface {
	Produce(msg entity.Message) error
	Close()
}
