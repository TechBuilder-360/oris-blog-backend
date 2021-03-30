package config

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//GetEntityDbCollection get collection from database
func GetEntityDbCollection(DbName string, CollectionName string) (*mongo.Collection) {

	URI := os.Getenv("DATABASE_ADDRESS")
	// URI := "mongodb://localhost:27017"
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(URI))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database(DbName).Collection(CollectionName)
	
	return collection
}