package services

import (
	"github.com/go-pg/pg"
	"github.com/webcat12345/go-one/core/repository"
	"github.com/webcat12345/go-one/model"
)

type (
	UserService interface {
		GetUsers() ([]*model.User, error)
	}
	DefaultUserService struct {
		userRepository repository.UserRepository
	}
)

func NewUserService(db *pg.DB) UserService {
	return &DefaultUserService{
		userRepository: repository.NewUserRepository(db),
	}
}

func (s *DefaultUserService) GetUsers() ([]*model.User, error) {
	return s.userRepository.FindAll()
}
