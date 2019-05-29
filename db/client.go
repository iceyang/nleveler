package db

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database *mongo.Database

func Connect() error {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		return err
	}

	// Check the connection
	if err = client.Ping(ctx, nil); err != nil {
		return err
	}

	database = client.Database("test")
	fmt.Println("Connected to MongoDB!")
	return nil
}

func Collection(name string) (*mongo.Collection, error) {
	if database == nil {
		return nil, errors.New("mongodb is not ready.")
	}
	return database.Collection(name), nil
}
