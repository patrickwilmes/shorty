package surl

import (
	"crypto/sha1"
	"encoding/base64"
	"github.com/patrickwilmes/shorty/internal/models"
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

func (su surlService) Create(targeturl string, token models.Token) (string, error) {
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
