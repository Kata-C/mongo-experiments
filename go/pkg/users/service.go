package users

import (
	"log"

	"github.com/mongo-experiments/go/pkg/api"
)

type storage interface {
	GetUsers() ([]api.User, error)
	CreateUser(user api.User) (*api.User, error)
}

type Service struct {
	storage storage
}

func NewService(storage storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s Service) GetUsers() ([]api.User, error) {
	users, err := s.storage.GetUsers()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return users, nil
}

func (s Service) CreateUser(user api.User) (*api.User, error) {
	dbUser, err := s.storage.CreateUser(user)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return dbUser, nil
}
