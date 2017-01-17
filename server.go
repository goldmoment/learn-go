package main

import (
	"github.com/goldmoment/manager"
	"github.com/goldmoment/router"
)

func main() {
	db.InitDatabase()
	router.InitServer()
}
