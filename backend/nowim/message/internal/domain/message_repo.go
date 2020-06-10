package domain

type MessageRepo interface {
	InsertMessage(message *Message) error
	QueryMessage(senderID, receiverID, limit int64) ([]*Message, error)
}
