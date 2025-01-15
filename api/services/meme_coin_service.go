package services

import (
	"context"
	"errors"
	"time"

	"github.com/cheildo/meme_coin_api/api/models"
	"github.com/cheildo/meme_coin_api/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "meme_coins"

func CreateMemeCoin(coin models.MemeCoin) (models.MemeCoin, error) {
	collection := config.GetCollection(collectionName)

	coin.CreatedAt = time.Now()
	coin.PopularityScore = 0

	_, err := collection.InsertOne(context.TODO(), coin)
	if mongo.IsDuplicateKeyError(err) {
		return models.MemeCoin{}, errors.New("a meme coin with this name already exists")
	} else if err != nil {
		return models.MemeCoin{}, err
	}
	return coin, nil
}

func GetMemeCoin(id string) (models.MemeCoin, error) {
	collection := config.GetCollection(collectionName)
	var coin models.MemeCoin

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return coin, err
	}

	err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&coin)
	if err != nil {
		return coin, err
	}
	return coin, nil
}

func UpdateDescription(id string, description string) error {
	collection := config.GetCollection(collectionName)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"description": description}}

	_, err = collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func DeleteMemeCoin(id string) error {
	collection := config.GetCollection(collectionName)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	return err
}

func PokeMemeCoin(id string) error {
	collection := config.GetCollection(collectionName)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$inc": bson.M{"popularity_score": 1}}

	_, err = collection.UpdateOne(context.TODO(), filter, update)
	return err
}
