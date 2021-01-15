package tokens

import (
	"errors"
	"github.com/google/uuid"
	"github.com/patrickwilmes/shorty/internal/models"
)

const (
	emptyToken = models.Token("")
)

var (
	errTokenAlreadyExists = errors.New("the generated token already exists")
)

type Repository interface {
	Create(token models.Token) error
	Delete(token models.Token) error
	Exists(token models.Token) (bool, error)
}

type Service interface {
	Create() (models.Token, error)
	Delete(token models.Token) error
}

type srv struct {
	repository Repository
}

func New(repository Repository) Service {
	return &srv{
		repository: repository,
	}
}

func (s srv) Create() (models.Token, error) {
	generatedUuid, _ := uuid.NewRandom()
	token := models.Token(generatedUuid.String())
	exists, err := s.repository.Exists(token)
	if err != nil {
		return emptyToken, err
	}
	if exists {
		return emptyToken, errTokenAlreadyExists
	}
	return token, s.repository.Create(token)
}

func (s srv) Delete(token models.Token) error {
	return s.repository.Delete(token)
}
