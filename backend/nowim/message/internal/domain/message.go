package domain

import (
	"nowim.message/pkg/channel"
	"time"

	"nowim.message/pkg/ulid"
)

type Message struct {
	MessageID  string `json:"messageID" bson:"messageID"`
	SenderID   int64  `json:"senderID" bson:"senderID"`
	ReceiverID int64  `json:"receiverID" bson:"receiverID"`
	Content    string `json:"content" bson:"content"`
	ClientTime int64  `json:"clientTime" bson:"clientTime"`
	ServerTime int64  `json:"serverTime" bson:"serverTime"`
}

func NewMessage(senderID, receiverID, clientTime int64, content string) *Message {
	return &Message{
		MessageID:  ulid.ULID(),
		SenderID:   senderID,
		ReceiverID: receiverID,
		Content:    content,
		ClientTime: clientTime,
		ServerTime: time.Now().UnixNano(),
	}
}

func (m *Message) ChannelMessage() *channel.Message {
	return &channel.Message{
		MessageID:  m.MessageID,
		SenderID:   m.SenderID,
		ReceiverID: m.ReceiverID,
		Content:    m.Content,
		ClientTime: m.ClientTime,
		ServerTime: m.ServerTime,
	}
}
