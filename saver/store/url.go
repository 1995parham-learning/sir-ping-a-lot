package store

import (
	"log"
	"saver/model"

	"gorm.io/gorm"
)

type URL interface {
	Insert(url model.URL) error
	GetTable() ([]model.URL, error)
}

type SQLURL struct {
	DB *gorm.DB
}

func NewURL(d *gorm.DB) SQLURL {
	return SQLURL{DB: d}
}

// Creates a table in the database that matches the URL table.
func (u SQLURL) Create() {
	//if err := u.DB.Migrator().DropTable(&model.URL{}); err != nil {
	//	log.Fatal(err)
	//}

	if err := u.DB.Migrator().CreateTable(&model.URL{}); err != nil {
		log.Fatal(err)
	}
}
