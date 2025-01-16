package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/cheildo/meme_coin_api/api/routes"
	"github.com/cheildo/meme_coin_api/config"
)

func main() {
	uri := os.Getenv("MONGO_URI")
	fmt.Println(uri)
	dbName := "meme_coin_db"
	collectionName := "meme_coins"

	_, err := config.ConnectDB(uri, dbName, collectionName)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	router := routes.SetupRouter()
	router.Run(":8080")
}
