package db

import "go.mongodb.org/mongo-driver/mongo"

const (
	tokenCollection = "tokens"
	shortUrlCollection = "short_urls"
	databaseName = "shorty"
)

func getCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database(databaseName).Collection(collectionName)
}
