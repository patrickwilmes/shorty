package handler

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)


type tokenContext struct {
	client *mongo.Client
}

func InitializeTokenHandlers(router *mux.Router, client *mongo.Client) {
	handlerContext := tokenContext{client: client}
	router.Methods(methodPost).Path("/token").Name("CreateToken").HandlerFunc(handlerContext.createToken)
}

func (tc tokenContext) createToken(w http.ResponseWriter, r *http.Request) {}
