package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/patrickwilmes/shorty/internal/handler"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

const (
	configServerAddress = "server.address"
	configDatabaseConnectionUrl = "database.url"
)

func main() {
	initializeConfiguration()
	startServer()
}

func initializeConfiguration() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
}

func initializeDatabase() (*mongo.Client, context.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI(viper.GetString(configDatabaseConnectionUrl)))
	if err != nil {
		log.Fatalln(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return client, ctx
}

func configureRouter(client *mongo.Client) *mux.Router {
	router := mux.NewRouter()
	handler.InitializeTokenHandlers(router, client)
	handler.InitializeShortUrlHandlers(router, client)
	handler.InitializeRedirectHandlers(router, client)
	return router
}

func startServer() {
	client, ctx := initializeDatabase()
	defer func() {
		log.Println("Disconnecting from database!")
		err := client.Disconnect(ctx)
		if err != nil {
			panic(err)
		}
	}()
	serverAddress := viper.GetString(configServerAddress)
	log.Printf("Starting server at %s\n", serverAddress)
	err := http.ListenAndServe(serverAddress, configureRouter(client))
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Shutting down server!\n")
}
