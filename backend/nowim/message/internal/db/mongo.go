package db

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"nowim.message/internal/config"
	"time"
)

var client *mongo.Client

func init() {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	uri := fmt.Sprintf("mongodb://%s:%s", config.Config().Mongo.Host, config.Config().Mongo.Port)

	var err error
	if client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri)); err != nil {
		log.Fatalf("can't connect to mongo, err: %+v", err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("ping mongo failed, err: %+v", err)
	}
	log.Info("mongo connected")
}

func MongoCollection(database, collection string) *mongo.Collection {
	log.Infof("get mongo collection, database: %s, collection: %s", database, collection)
	return client.Database(database).Collection(collection)
}
