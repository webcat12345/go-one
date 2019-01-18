package main

import "github.com/webcat12345/go-one/route"

func main() {
	e := route.Init()
	e.Logger.Fatal(e.Start(":1323"))
}
