package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func DBInstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: Unable to load .env file")
	}

	MongoDBURI := os.Getenv("MONGODB_URI")
	if MongoDBURI == "" {
		log.Fatal("Warning: MONGODB_URI is not set")
	}
	fmt.Println("MongoDBURI:", MongoDBURI)
	clientOptions := options.Client().ApplyURI(MongoDBURI)
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	return client
}
