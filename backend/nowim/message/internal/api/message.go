package api

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
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
	return messageServerImpl{
		repo: domain.NewMemMessageRepo(),
		//messageChan TODO
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
