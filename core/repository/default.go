package repository

import (
	"github.com/go-pg/pg"
)

func GetDatabase() (*pg.DB, error) {
	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5433",
		User:     "postgres",
		Password: "password",
		Database: "go_one",
	})
	// execute test query to check connection status
	_, err := db.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}
	return db, nil
}
