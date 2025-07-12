package config

import (
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var DbClient *mongo.Database

func InitDB() (*mongo.Client, error) {
	log.Println("Initializing MongoDB connection...")

	// client options
	clientOptions := options.Client().ApplyURI(Constants.MongoURI)

	// connect to MongoDB
	client, err := mongo.Connect(clientOptions)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Println("Connected to MongoDB!")

	DbClient = client.Database(Constants.MongoDBName)

	return client, nil
}
