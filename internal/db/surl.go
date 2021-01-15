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

func (su ShortUrlRepository) GetByToken(token models.Token) (resultingModels []models.ShortUrl, err error) {
	collection := getCollection(su.Client, shortUrlCollection)
	result, err := collection.Find(context.Background(), bson.M{"token": token})
	if err != nil {
		return nil, err
	}
	defer result.Close(context.Background())
	for result.Next(context.Background()) {
		var url models.ShortUrl
		err = result.Decode(&url)
		if err != nil {
			return nil, err
		}
		resultingModels = append(resultingModels, url)
	}
	return
}
func (su ShortUrlRepository) GetByHash(hash string) (models.ShortUrl, error) {
	collection := getCollection(su.Client, shortUrlCollection)
	result := collection.FindOne(context.Background(), bson.M{"hash": hash})
	var model models.ShortUrl
	err := result.Decode(&model)
	return model, err
}
