package handler

import (
	"github.com/gorilla/mux"
	"github.com/patrickwilmes/shorty/internal/db"
	"github.com/patrickwilmes/shorty/internal/redirect"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type redirectContext struct {
	client  *mongo.Client
	service redirect.Service
}

func InitializeRedirectHandlers(router *mux.Router, client *mongo.Client) {
	repo := db.ShortUrlRepository{Client: client}
	handlerContext := redirectContext{client: client, service: redirect.New(repo)}
	router.Methods(methodGet).Path("/{hash}").Name("Redirect").HandlerFunc(handlerContext.redirectTo)
}

func (rd redirectContext) redirectTo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash := vars["hash"]
	targetUrl, err := rd.service.GetTargetUrlForHash(hash)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, targetUrl, 301)
}
