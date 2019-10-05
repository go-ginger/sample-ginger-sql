package models

import (
	"github.com/go-ginger/sql"
	"time"
)

type Movie struct {
	sql.BaseModel

	Title       string    `json:"title,omitempty"`
	ReleaseDate time.Time `json:"release_date,omitempty"`
	Genres      []Genre   `json:"genres,omitempty" gorm:"many2many:movie_genres;"`
	Synopsis    string    `json:"synopsis,omitempty" gorm:"type:TEXT"`
}
