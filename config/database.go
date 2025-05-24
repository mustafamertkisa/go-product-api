package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var ProductCollection *mongo.Collection

func Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	LoadConfig()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(AppConfig.MongoURI))
	if err != nil {
		log.Fatalf("mongodb connection failed: %v", err)
	}

	MongoClient = client
	ProductCollection = client.Database("test").Collection("products")
}
