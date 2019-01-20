package services

import (
	"github.com/go-pg/pg"
	"github.com/labstack/echo"
	"github.com/webcat12345/go-one/core/repository"
	"github.com/webcat12345/go-one/core/tokenizer"
	"github.com/webcat12345/go-one/model"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type (
	UserService interface {
		GetUsers() ([]*model.User, error)
		GetUserById(id int) (*model.User, error)
		CreateUser(email, password string) (*model.User, error)
		Login(email, password string) (map[string]string, error)
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

func (s *DefaultUserService) Login(email, password string) (map[string]string, error) {
	// check user existance
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "User does not exists")
	}
	// check password
	if err := s.ComparePassword(user.Password, password); err != nil {
		return nil, err
	}
	// build access token
	token, err := tokenizer.NewAccessToken(user.Id)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Failed to create jwt token")
	}

	return map[string]string{"token": token}, nil
}

func (s *DefaultUserService) ComparePassword(hash []byte, password string) error {
	if err := bcrypt.CompareHashAndPassword(hash, []byte(password)); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Password does not match")
	}
	return nil
}
