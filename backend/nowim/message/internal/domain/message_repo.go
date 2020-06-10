package domain

type MessageRepo interface {
	SaveMessage(message *Message) error
	QueryMessage(senderID, receiverID, limit int64) ([]*Message, error)
}
