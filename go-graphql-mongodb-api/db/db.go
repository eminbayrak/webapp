package db

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func InitMongoDB() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoURL := os.Getenv("MONGODB_URL")

	clientOptions := options.Client().ApplyURI(mongoURL)
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
}
