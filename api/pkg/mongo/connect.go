package mongo

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	// TODO: correctly handle connection error
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")))

	if err != nil {
		return nil, err
	}

	return client, nil
}
