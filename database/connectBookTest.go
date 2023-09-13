package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func FakeConnectBook() *mongo.Collection {
	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("connection-key").SetServerAPIOptions(serverApi)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}
	databases := client.Database("apiv0Test")
	bookCollection := databases.Collection("booksTest")

	return bookCollection
}
