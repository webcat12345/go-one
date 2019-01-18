package main

import (
	"github.com/webcat12345/go-one/core/repository"
	"github.com/webcat12345/go-one/route"
	"log"
)

func main() {

	// create db connection
	db, err := repository.GetDatabase()
	if err != nil {
		log.Fatal(err)
	}

	e := route.Init(db)
	e.Logger.Fatal(e.Start(":1323"))

	// close db connection
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}
