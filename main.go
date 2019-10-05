package main

import (
	"github.com/go-ginger/sample-ginger-sql/project/controllers"
	"github.com/go-ginger/sample-ginger-sql/project/db"
	"github.com/go-ginger/sample-ginger-sql/project/global"
	"log"
)

func main() {
	global.Initialize()
	db.Initialize()
	engine := controllers.RegisterRoutes()
	err := engine.Run()
	log.Fatal(err)
}
