package db

import (
	"context"
	"fmt"
	"github.com/iarsham/task-realtime-app/chat-service/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func OpenDB(cfg *configs.Config) (*mongo.Client, error) {
	dsn := makeDsn(cfg)
	clientOpts := options.Client().ApplyURI(dsn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	return client, nil
}

func makeDsn(cfg *configs.Config) string {
	return fmt.Sprintf(
		"mongodb://%s:%s@%s:%d",
		cfg.Mongo.MongoUser,
		cfg.Mongo.MongoPass,
		cfg.Mongo.MongoHost,
		cfg.Mongo.MongoPort,
	)
}
