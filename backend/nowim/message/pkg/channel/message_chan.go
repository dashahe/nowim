package channel

import (
	"nowim.message/pkg/ulid"
	"time"
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

type MessageChannel interface {
	PushMessage(channel int64, message *Message) error
	Consume(channel int64) (*Message, error)
	Close(channel int64) error
}
