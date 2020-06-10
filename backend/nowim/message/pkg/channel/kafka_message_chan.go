package channel

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"nowim.message/internal/domain"
	"strconv"
)

type KafkaMessageChannel struct {
	writers map[int64]*kafka.Writer
	readers map[int64]*kafka.Reader
	writerConfig kafka.WriterConfig
	readerConfig kafka.ReaderConfig
}

func (k KafkaMessageChannel) PushMessage(channel int64, message *domain.Message) error {
	if k.writers[channel] == nil {
		k.writerConfig.Topic = strconv.FormatInt(channel, 10)
		k.writers[channel] = kafka.NewWriter(k.writerConfig)
	}

	data, _ := json.Marshal(message)
	kafkaMessage := kafka.Message{
		Key:       []byte(message.MessageID),
		Value:     data,
	}
	return k.writers[channel].WriteMessages(context.Background(), kafkaMessage)
}

// TODO，考虑close的signal提醒，避免websocket连接断开后，继续消费消息
func (k KafkaMessageChannel) Consume(channel int64) (*domain.Message, error) {
	if k.readers[channel] == nil {
		k.readerConfig.Topic = strconv.FormatInt(channel, 10)
		k.readers[channel] = kafka.NewReader(k.readerConfig)
	}

	kafkaMessage, err := k.readers[channel].ReadMessage(context.Background())
	if err != nil {
		return nil, err
	}

	var message domain.Message
	if err := json.Unmarshal(kafkaMessage.Value, &message); err != nil {
		return nil, err
	}

	return &message, nil
}
