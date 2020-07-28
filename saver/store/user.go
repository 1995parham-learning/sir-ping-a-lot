package store

import (
	"database/sql"
	"log"
	"saver/config"
	"saver/model"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
func (u SQLUser) Create(config config.Database) {
	db, err := sql.Open("postgres", config.Cstring())
	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	p, err := migrate.NewWithDatabaseInstance("file://./migration", "monitor", driver)
	if err != nil {
		log.Fatal(err)
	}

	if err := p.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
