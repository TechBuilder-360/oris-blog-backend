package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//GetEntityDbCollection get collection from database
func GetEntityDbCollection(DbName string, DbAddress string, CollectionName string) (*mongo.Collection) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(DbAddress))

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