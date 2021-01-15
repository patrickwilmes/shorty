package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/patrickwilmes/shorty/internal/db"
	"github.com/patrickwilmes/shorty/internal/models"
	"github.com/patrickwilmes/shorty/internal/surl"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type surlContext struct {
	client  *mongo.Client
	service surl.Service
}

type shortUrlDto struct {
	TargetUrl string
	Token     models.Token
}

func InitializeShortUrlHandlers(router *mux.Router, client *mongo.Client) {
	repo := db.ShortUrlRepository{Client: client}
	handlerContext := surlContext{client: client, service: surl.New(repo)}
	router.Methods(methodPost).Path("/url").Name("CreateShortUrl").HandlerFunc(handlerContext.createShortUrl)
	router.Methods(methodDelete).Path("/url/{shortUrlId}/{token}").Name("DeleteShortUrl").HandlerFunc(handlerContext.deleteShortUrl)
}

// todo - implement proper error handling with some kind of problem json
func (sc surlContext) createShortUrl(w http.ResponseWriter, r *http.Request) {
	var shortUrlDto shortUrlDto
	err := json.NewDecoder(r.Body).Decode(&shortUrlDto)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := sc.service.Create(shortUrlDto.TargetUrl, shortUrlDto.Token)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_ = json.NewEncoder(w).Encode(struct {
		ShortUrlId string
	}{
		ShortUrlId: id,
	})
}

func (sc surlContext) deleteShortUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token := models.Token(vars["token"])
	shortUrlId := vars["shortUrlId"]
	err := sc.service.Delete(token, shortUrlId)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
