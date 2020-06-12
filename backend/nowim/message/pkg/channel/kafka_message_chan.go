package channel

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/segmentio/kafka-go"
)

type KafkaMessageChannel struct {
	writers      map[int64]*kafka.Writer
	readers      map[int64]*kafka.Reader
	writerConfig kafka.WriterConfig
	readerConfig kafka.ReaderConfig
}

func NewKafkaMessageChannel(writerConfig kafka.WriterConfig, readerConfig kafka.ReaderConfig) *KafkaMessageChannel {
	return &KafkaMessageChannel{
		writers:      make(map[int64]*kafka.Writer),
		readers:      make(map[int64]*kafka.Reader),
		writerConfig: writerConfig,
		readerConfig: readerConfig,
	}
}

func (k KafkaMessageChannel) PushMessage(channel int64, message *Message) error {
	if k.writers[channel] == nil {
		k.writerConfig.Topic = strconv.FormatInt(channel, 10)
		k.writers[channel] = kafka.NewWriter(k.writerConfig)
	}

	data, _ := json.Marshal(message)
	kafkaMessage := kafka.Message{
		Key:   []byte(message.MessageID),
		Value: data,
	}
	return k.writers[channel].WriteMessages(context.Background(), kafkaMessage)
}

func (k KafkaMessageChannel) Consume(channel int64) (*Message, error) {
	if k.readers[channel] == nil {
		k.readerConfig.Topic = strconv.FormatInt(channel, 10)
		k.readers[channel] = kafka.NewReader(k.readerConfig)
	}

	kafkaMessage, err := k.readers[channel].ReadMessage(context.Background())
	if err != nil {
		return nil, err
	}

	var message Message
	if err := json.Unmarshal(kafkaMessage.Value, &message); err != nil {
		return nil, err
	}

	return &message, nil
}

func (k KafkaMessageChannel) Close(channel int64) error {
	if r, ok := k.readers[channel]; ok {
		if err := r.Close(); err != nil {
			return err
		}
	}
	if w, ok := k.writers[channel]; ok {
		if err := w.Close(); err != nil {
			return err
		}
	}
	delete(k.readers, channel)
	delete(k.writers, channel)
	return nil
}
