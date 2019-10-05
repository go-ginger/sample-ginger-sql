package models

import (
	"github.com/go-ginger/sql"
)

type Genre struct {
	sql.BaseModel

	Name string `json:"name,omitempty" gorm:"type:varchar(24)" sql:"index"`
}
