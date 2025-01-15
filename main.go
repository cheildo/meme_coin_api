package main

import (
	"log"

	"github.com/cheildo/meme_coin_api/api/routes"
	"github.com/cheildo/meme_coin_api/config"
)

func main() {
	uri := "mongodb://localhost:27017"
	dbName := "meme_coin_db"
	collectionName := "meme_coins"

	_, err := config.ConnectDB(uri, dbName, collectionName)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	router := routes.SetupRouter()
	router.Run(":8080")
}
