package services

import (
	"github.com/go-pg/pg"
	"github.com/labstack/echo"
	"github.com/webcat12345/go-one/core/repository"
	"github.com/webcat12345/go-one/model"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type (
	UserService interface {
		GetUsers() ([]*model.User, error)
		GetUserById(id int) (*model.User, error)
		CreateUser(email, password string) (*model.User, error)
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
	res, err := s.userRepository.FindAll()
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Failed to get users")
	}
	return res, nil
}

func (s *DefaultUserService) GetUserById(id int) (*model.User, error) {
	res, err := s.userRepository.FindById(id)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Failed to get user")
	}
	return res, nil
}

func (s *DefaultUserService) CreateUser(email, password string) (*model.User, error) {

	// check if email already exists
	if s.userRepository.ExistsByEmail(email) {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Email already exists")
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to hashing a password")
	}

	user := &model.User{
		Email:    email,
		Password: hash,
	}
	return s.userRepository.Create(user)
}
