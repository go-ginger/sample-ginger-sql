package controllers

import (
	g "github.com/go-ginger/ginger"
	v1 "github.com/go-ginger/sample-ginger-sql/project/controllers/v1"
	"github.com/go-ginger/sample-ginger-sql/project/global"
)

func RegisterRoutes() *g.Router {
	router := g.NewRouter()
	router.Address = []string{"0.0.0.0:" + global.Config.Port}
	routerV1 := router.Group("/v1")
	v1.RegisterRoutes(routerV1)
	return router
}
