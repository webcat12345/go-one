package repository

import (
	"github.com/go-pg/pg"
)

func GetDatabase() *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5433",
		User:     "postgres",
		Password: "password",
		Database: "go_one",
	})
	return db
}
