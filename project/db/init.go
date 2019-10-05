package db

import (
	"fmt"
	"github.com/go-ginger/sample-ginger-sql/project/models"
	"github.com/go-ginger/sql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var MovieDbHandler = movieSqlHandler{}
var GenreDbHandler = genreSqlHandler{}

func Initialize() {
	db, err := sql.GetDb()
	if err != nil {
		log.Println(fmt.Sprintf("Error on connect to database, Error: %v", err))
		panic(err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Println(fmt.Sprintf("Error while closing database, Error: %v", err))
			panic(err)
		}
	}()
	db = db.Set("gorm:table_options", " ENGINE=InnoDB CHARSET=utf8mb4")

	movie := models.Movie{}
	genre := models.Genre{}
	movieGenre := models.MovieGenre{}

	db.AutoMigrate(&movie)
	db.AutoMigrate(&genre)
	db.AutoMigrate(&movieGenre)

	db.Model(&movieGenre).AddForeignKey("movie_id", "movies(id)", "NO ACTION", "NO ACTION")
	db.Model(&movieGenre).AddForeignKey("genre_id", "genres(id)", "NO ACTION", "NO ACTION")
}
