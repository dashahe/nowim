package api

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"nowim.message/internal/config"
	"nowim.message/internal/domain"
	"nowim.message/pkg/channel"
	"nowim.message/pkg/message"
	"nowim.message/pkg/ulid"
	"time"
)

type messageServerImpl struct {
	message.UnimplementedMessageServer
	repo        domain.MessageRepo
	messageChan channel.MessageChannel
}

func NewMessageServer() message.MessageServer {
	w := kafka.WriterConfig{
		Brokers: config.Config().Kafka.Brokers,
		Balancer: &kafka.LeastBytes{},
	}
	r := kafka.ReaderConfig{
		Brokers: config.Config().Kafka.Brokers,
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	}
	return messageServerImpl{
		repo: domain.NewMongoMessageRepo(),
		messageChan: channel.NewKafkaMessageChannel(w, r),
	}
}

func (m messageServerImpl) SendMessage(ctx context.Context, req *message.SendMessageRequest) (*message.SendMessageReply, error) {
	newMsg := domain.Message{
		MessageID:  ulid.ULID(),
		SenderID:   req.SenderID,
		ReceiverID: req.ReceiverID,
		Content:    req.Content,
		ClientTime: req.ClientTime,
		ServerTime: time.Now().UnixNano(),
	}

	// 存入mongo db
	if err := m.repo.InsertMessage(&newMsg); err != nil {
		log.Errorf("save message failed, err: %+v", err)
		return nil, fmt.Errorf("save message failed, err: %+v", err)
	}

	// push到对应的message channel
	if err := m.messageChan.PushMessage(newMsg.ReceiverID, &newMsg); err != nil {
		log.Errorf("push message to message channel failed, err: %+v", err)
		return nil, fmt.Errorf("push message to message channel failed, err: %+v", err)
	}

	return nil, nil
}
