package database

import (
	"context"
	"log"
	// "os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//GetMongoDbCollection get collection from mongodb
func GetMongoDbCollection(DbName string, CollectionName string) (*mongo.Collection) {

	// URI := os.Getenv("DATABASE_ADDRESS")
	URI := "mongodb://localhost:27017"
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