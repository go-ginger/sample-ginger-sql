package logics

import (
	l "github.com/go-ginger/logic"
	"github.com/go-ginger/sample-ginger-sql/project/db"
)

var MovieLogicHandler movieLogic
var GenreLogicHandler genreLogic

func init()  {
	MovieLogicHandler = movieLogic{IBaseLogicHandler: &l.BaseLogicHandler{}}
	GenreLogicHandler = genreLogic{IBaseLogicHandler: &l.BaseLogicHandler{}}

	MovieLogicHandler.Init(&MovieLogicHandler, &db.MovieDbHandler)
	GenreLogicHandler.Init(&GenreLogicHandler, &db.GenreDbHandler)
}
