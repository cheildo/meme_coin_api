package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB(uri string, dbName string, collectionName string) (*mongo.Collection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	DB = client.Database("meme_coin_db")
	log.Println("Connected to MongoDB!")

	collection := DB.Collection(collectionName)

	// Ensure unique index on "name"
	if err := CreateUniqueIndex(collection); err != nil {
		log.Fatalf("failed to create unique index on name: %v", err)
	}

	return collection, nil
}

func GetCollection(collectionName string) *mongo.Collection {
	return DB.Collection(collectionName)
}

func CreateUniqueIndex(collection *mongo.Collection) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	indexModel := mongo.IndexModel{
		Keys:    bson.M{"name": 1}, // Index on the "name" field
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	return err
}
