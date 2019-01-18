package repository

import (
	"github.com/go-pg/pg"
	"github.com/webcat12345/go-one/model"
)

type (
	UserRepository interface {
		FindAll() ([]*model.User, error)
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
	} else {
		return users, nil
	}
}
