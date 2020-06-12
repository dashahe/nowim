package channel

type MemMessageChannel struct {
	channels map[int64]chan *Message
}

func NewMemMessageChannel() *MemMessageChannel {
	return &MemMessageChannel{
		channels: make(map[int64]chan *Message),
	}
}

func (m *MemMessageChannel) PushMessage(channel int64, message *Message) error {
	m.checkEmptyChan(channel)
	m.channels[channel] <- message
	return nil
}

func (m *MemMessageChannel) Consume(channel int64) (*Message, error) {
	m.checkEmptyChan(channel)
	result := <-m.channels[channel]
	return result, nil
}

func (m *MemMessageChannel) Close(channel int64) error {
	delete(m.channels, channel)
	return nil
}

func (m *MemMessageChannel) checkEmptyChan(channel int64) {
	_, ok := m.channels[channel]
	if !ok {
		m.channels[channel] = make(chan *Message, 1024)
	}
}
