package db

import (
	"context"
	"github.com/patrickwilmes/shorty/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TokenDto struct {
	Token string
}

type TokenRepository struct {
	Client *mongo.Client
}

func (tr TokenRepository) Create(token models.Token) error {
	_, err := getCollection(tr.Client, tokenCollection).InsertOne(context.Background(), TokenDto{Token: string(token)})
	return err
}

func (tr TokenRepository) Delete(token models.Token) error {
	_, err := getCollection(tr.Client, tokenCollection).DeleteOne(context.Background(), bson.M{"token": token})
	return err
}

func (tr TokenRepository) Exists(token models.Token) (bool, error) {
	result := getCollection(tr.Client, tokenCollection).FindOne(context.Background(), bson.M{"token": token})
	var tokenDto TokenDto
	err := result.Decode(&tokenDto)
	return err == nil, nil // todo - this error could be removed here as FindOne is only returning a result and no error
}
