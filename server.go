package main

import (
	"./manager"
	"./router"
)

func main() {
	db.InitDatabase()
	router.InitServer()
}
