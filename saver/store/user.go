package store

import (
	"log"
	"saver/model"

	"gorm.io/gorm"
)

type User interface {
	Insert(user model.User) error
	Retrieve(user model.User) (model.User, error)
}

type SQLUser struct {
	DB *gorm.DB
}

func NewUser(d *gorm.DB) SQLUser {
	return SQLUser{DB: d}
}

// Creates a table in the database that matches the User table.
func (u SQLUser) Create() {
	if err := u.DB.Migrator().DropTable(&model.User{}); err != nil {
		log.Fatal(err)
	}

	if err := u.DB.Migrator().CreateTable(&model.User{}); err != nil {
		log.Fatal(err)
	}
}