package config

import (
	"context"
	"go-crud/app/helper"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDatabase(configuration Config) *mongo.Database {
	ctx, cancel := NewMongoContext()
	defer cancel()

	mongoPoolMin, err := strconv.Atoi(configuration.Get("MONGO_POOL_MIN", "10"))
	helper.FatalIfNeeded(err)

	mongoPoolMax, err := strconv.Atoi(configuration.Get("MONGO_POOL_MAX", "100"))
	helper.FatalIfNeeded(err)

	mongoMaxIdleTime, err := strconv.Atoi(configuration.Get("MONGO_MAX_IDLE_TIME_SECOND", "60"))
	helper.FatalIfNeeded(err)

	option := options.Client().
		ApplyURI(configuration.Get("MONGO_URI", "mongodb://mongo:mongo@localhost:27017")).
		SetMinPoolSize(uint64(mongoPoolMin)).
		SetMaxPoolSize(uint64(mongoPoolMax)).
		SetMaxConnIdleTime(time.Duration(mongoMaxIdleTime) * time.Second)

	client, err := mongo.NewClient(option)
	helper.FatalIfNeeded(err)

	err = client.Connect(ctx)
	helper.FatalIfNeeded(err)

	database := client.Database(configuration.Get("MONGO_DATABASE", "go-crud"))
	return database
}

func NewMongoContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
