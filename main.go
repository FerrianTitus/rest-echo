package main

import (
	"github.com/ferriantitus/rest-echo/db"
	"github.com/ferriantitus/rest-echo/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":2121"))
}