package db

import (
	"context"
	"github.com/patrickwilmes/shorty/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ShortUrlRepository struct {
	Client *mongo.Client
}

func (su ShortUrlRepository) Create(url models.ShortUrl) error {
	collection := getCollection(su.Client, shortUrlCollection)
	_, err := collection.InsertOne(context.Background(), url)
	return err
}
func (su ShortUrlRepository) Delete(token models.Token, id string) error {
	collection := getCollection(su.Client, shortUrlCollection)
	_, err := collection.DeleteOne(context.Background(), bson.M{"id": id, "token": token})
	return err
}
