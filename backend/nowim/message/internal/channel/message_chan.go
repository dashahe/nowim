package channel

import "nowim.message/internal/domain"

type MessageChannel interface {
	PushMessage(message *domain.Message) error
}
