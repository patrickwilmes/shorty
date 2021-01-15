package redirect

import "github.com/patrickwilmes/shorty/internal/db"

type Service interface {
	GetTargetUrlForHash(hash string) (string, error)
}

type redirectService struct {
	ShortUrlRepository db.ShortUrlRepository
}

func New(shortUrlRepository db.ShortUrlRepository) Service {
	return &redirectService{ShortUrlRepository: shortUrlRepository}
}

func (rs redirectService) GetTargetUrlForHash(hash string) (string, error) {
	shortUrl, err := rs.ShortUrlRepository.GetByHash(hash)
	if err != nil {
		return "", err
	}
	return shortUrl.TargetUrl, nil
}
