package main

import (
	"github.com/goldmoment/learn-go/manager"
	"github.com/goldmoment/learn-go/router"
)

func main() {
	db.InitDatabase()
	router.InitServer()
}
