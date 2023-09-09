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
	opts := options.Client().ApplyURI("mongodb+srv://yigithankarabulut:Yigitk12..12.@cluster0.wqmtmvm.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverApi)

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
