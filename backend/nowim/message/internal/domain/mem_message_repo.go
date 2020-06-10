package domain

import "sync"

type MemMessageRepo struct {
	mu   sync.RWMutex
	data []*Message
}

func NewMemMessageRepo() *MemMessageRepo {
	return &MemMessageRepo{data: make([]*Message, 0)}
}

func (m *MemMessageRepo) SaveMessage(message *Message) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.data = append(m.data, message)
	return nil
}

func (m *MemMessageRepo) QueryMessage(senderID, receiverID, limit int64) ([]*Message, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	result := make([]*Message, 0)
	var cnt int64
	for i := len(m.data) - 1; i >= 0 && cnt < limit; i -= 1 {
		msg := m.data[i]
		if msg.SenderID == senderID && msg.ReceiverID == receiverID {
			result = append(result, msg)
			cnt += 1
		}
	}
	return result, nil
}
