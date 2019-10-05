package v1

import (
	"github.com/go-ginger/sample-ginger-sql/project/models"
)

type MoviesController struct {
	models.BaseItemsController
}

type MovieController struct {
	models.BaseItemController
}
