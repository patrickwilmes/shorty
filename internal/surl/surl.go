package surl

import (
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"github.com/patrickwilmes/shorty/internal/models"
	"net/http"
)

var (
	errTargetUrlDoesNotExist = errors.New("target url does not exist or is not valid")
)

type Repository interface {
	Create(url models.ShortUrl) error
}

type Service interface {
	Create(targetUrl string, token models.Token) (string, error)
}

type surlService struct {
	SUrlRepository Repository
}

func New(surlRepository Repository) Service {
	return &surlService{SUrlRepository: surlRepository}
}

func urlExists(targetUrl string) bool {
	res, err := http.Head(targetUrl)
	if err != nil {
		return false
	}
	return res.StatusCode == 200
}

func (su surlService) Create(targeturl string, token models.Token) (string, error) {
	if !urlExists(targeturl) {
		return "", errTargetUrlDoesNotExist
	}
	h := sha1.New()
	h.Write([]byte(targeturl))
	bs := base64.URLEncoding.EncodeToString(h.Sum(nil))
	shortHash := bs[:10]
	shortUrl := models.ShortUrl{
		ID:        bs,
		TargetUrl: targeturl,
		Hash:      shortHash,
		Token:     token,
	}
	err := su.SUrlRepository.Create(shortUrl)
	return shortUrl.ID, err
}
