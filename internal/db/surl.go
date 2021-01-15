package db

import (
	"context"
	"github.com/patrickwilmes/shorty/internal/models"
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
