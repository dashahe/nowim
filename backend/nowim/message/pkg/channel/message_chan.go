package channel

import "nowim.message/internal/domain"

type MessageChannel interface {
	PushMessage(channel int64, message *domain.Message) error
	Consume(channel int64) (*domain.Message, error)
	Close(channel int64) error
}
