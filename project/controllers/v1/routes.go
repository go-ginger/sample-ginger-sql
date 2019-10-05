package v1

import (
	g "github.com/go-ginger/ginger"
	"github.com/go-ginger/sample-ginger-sql/project/db"
	"github.com/go-ginger/sample-ginger-sql/project/logics"
)

var movies = new(MoviesController)
var movie = new(MovieController)

var genres = new(GenresController)
var genre = new(GenreController)

func init() {
	movies.Init(movies, &logics.MovieLogicHandler, &db.MovieDbHandler)
	movie.Init(movie, &logics.MovieLogicHandler, &db.MovieDbHandler)
}

func RegisterRoutes(router *g.RouterGroup) {
	movies.AddRoute("Post")
	movies.AddRoute("Get")

	movie.AddRoute("Get")
	movie.AddRoute("Put")
	movie.AddRoute("Delete")

	genres.AddRoute("Post")
	genres.AddRoute("Get")

	genre.AddRoute("Get")
	genre.AddRoute("Put")
	genre.AddRoute("Delete")

	movies.RegisterRoutes(movies, "/movies", router)
	movie.RegisterRoutes(movie, "/movies/:id", router)

	genres.RegisterRoutes(movies, "/genres", router)
	genre.RegisterRoutes(movie, "/genres/:id", router)
}
