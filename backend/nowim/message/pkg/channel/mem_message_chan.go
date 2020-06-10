package channel

import (
	"nowim.message/internal/domain"
)

type MemMessageChannel struct {
	channels map[int64]chan *domain.Message
}

func NewMemMessageChannel() *MemMessageChannel {
	return &MemMessageChannel{
		channels: make(map[int64]chan *domain.Message),
	}
}

func (m *MemMessageChannel) PushMessage(channel int64, message *domain.Message) error {
	m.checkEmptyChan(channel)
	m.channels[channel] <- message
	return nil
}

func (m *MemMessageChannel) Consume(channel int64) (*domain.Message, error) {
	m.checkEmptyChan(channel)
	result := <-m.channels[channel]
	return result, nil
}

func (m *MemMessageChannel) checkEmptyChan(channel int64) {
	_, ok := m.channels[channel]
	if !ok {
		m.channels[channel] = make(chan *domain.Message, 1024)
	}
}
