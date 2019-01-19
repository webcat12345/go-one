package repository

import (
	"github.com/go-pg/pg"
	"github.com/webcat12345/go-one/model"
)

type (
	UserRepository interface {
		ExistsByEmail(email string) bool
		FindAll() ([]*model.User, error)
		FindById(id int) (*model.User, error)
		FindByEmail(email string) (*model.User, error)
		Create(*model.User) (*model.User, error)
	}
	DefaultUserRepository struct {
		db *pg.DB
	}
)

func NewUserRepository(db *pg.DB) UserRepository {
	return &DefaultUserRepository{db: db}
}

func (r *DefaultUserRepository) FindAll() ([]*model.User, error) {
	var users []*model.User
	if err := r.db.Model(&users).Select(); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *DefaultUserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Model(&user).Where("email = ?", email).Select()
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *DefaultUserRepository) FindById(id int) (*model.User, error) {
	user := &model.User{Id: id}
	err := r.db.Select(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *DefaultUserRepository) ExistsByEmail(email string) bool {
	_, err := r.FindByEmail(email)
	if err != nil {
		return false
	}
	return true
}

func (r *DefaultUserRepository) Create(user *model.User) (*model.User, error) {
	_, err := r.db.Model(user).Returning("*").Insert()
	if err != nil {
		return nil, err
	}
	return user, nil
}
