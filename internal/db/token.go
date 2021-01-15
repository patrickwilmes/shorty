package db

import (
	"context"
	"github.com/patrickwilmes/shorty/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// todo - place this in constants or so
	tokenCollection = "tokens"
	databaseName = "shorty"
)

type TokenDto struct {
	Token string
}

type TokenRepository struct {
	Client *mongo.Client
}

func (tr TokenRepository) getCollection() *mongo.Collection {
	return tr.Client.Database(databaseName).Collection(tokenCollection)
}

func (tr TokenRepository) Create(token models.Token) error {
	_, err := tr.getCollection().InsertOne(context.Background(), TokenDto{Token: string(token)})
	return err
}

func (tr TokenRepository) Delete(token models.Token) error {
	_, err := tr.getCollection().DeleteOne(context.Background(), bson.M{"token": token})
	return err
}

func (tr TokenRepository) Exists(token models.Token) (bool, error) {
	result := tr.getCollection().FindOne(context.Background(), bson.M{"token": token})
	var tokenDto TokenDto
	err := result.Decode(&tokenDto)
	return err == nil, nil // todo - this error could be removed here as FindOne is only returning a result and no error
}
