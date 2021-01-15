package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/patrickwilmes/shorty/internal/db"
	"github.com/patrickwilmes/shorty/internal/models"
	"github.com/patrickwilmes/shorty/internal/tokens"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type tokenContext struct {
	client *mongo.Client
	service tokens.Service
}

func InitializeTokenHandlers(router *mux.Router, client *mongo.Client) {
	repo := db.TokenRepository{Client: client}
	handlerContext := tokenContext{client: client, service: tokens.New(repo)}
	router.Methods(methodPost).Path("/token").Name("CreateToken").HandlerFunc(handlerContext.createToken)
	router.Methods(methodDelete).Path("/token/{token}").Name("DeleteToken").HandlerFunc(handlerContext.deleteToken)
}

func (tc tokenContext) deleteToken(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token := models.Token(vars["token"])
	err := tc.service.Delete(token)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		_ = json.NewEncoder(w).Encode(struct {
			message string
		}{
			message: err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (tc tokenContext) createToken(w http.ResponseWriter, r *http.Request) {
	token, err := tc.service.Create()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(struct {
			message string
		}{
			message: err.Error(),
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// todo - handle encoding errors
	_ = json.NewEncoder(w).Encode(struct {
		Token models.Token
	}{
		Token: token,
	})
}
