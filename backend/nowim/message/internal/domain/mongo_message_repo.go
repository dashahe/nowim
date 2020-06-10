package domain

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"nowim.message/internal/config"
	"nowim.message/internal/db"
	"time"
)

type MongoMessageRepo struct {
	mongoColl *mongo.Collection
}

func NewMongoMessageRepo() *MongoMessageRepo {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	database, collection := config.Config().Mongo.Database, config.Config().Mongo.Collection
	mongoColl := db.MongoCollection(database, collection)

	indexModel := mongo.IndexModel{Keys: bson.D{{"senderID", 1}, {"receiverID", 1}}}
	if _, err := mongoColl.Indexes().CreateOne(ctx, indexModel); err != nil {
		log.Errorf("create mongo index failed, err: %+v", err)
	}

	return &MongoMessageRepo{mongoColl: mongoColl}
}

func (m MongoMessageRepo) InsertMessage(message *Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := m.mongoColl.InsertOne(ctx, message)
	if err != nil {
		return fmt.Errorf("save messge err: %+v", err)
	}
	return nil
}

func (m MongoMessageRepo) QueryMessage(senderID, receiverID, limit int64) ([]*Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := Message{
		SenderID:   senderID,
		ReceiverID: receiverID,
	}
	result := make([]*Message, 0)
	opts := options.Find().SetSort(bson.D{{"messageID", -1}}).SetLimit(limit)

	cursor, err := m.mongoColl.Find(ctx, filter, opts)
	if err != nil {
		log.Errorf("mongo find failed, err: %+v", err)
		return nil, err
	}

	if err := cursor.All(ctx, &result); err != nil {
		log.Errorf("cursor get all failed, err: %+v", err)
		return nil, err
	}
	return result, nil
}
