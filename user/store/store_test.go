package store_test

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
	"user/config"
	"user/db"
	"user/model"
	"user/store"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	cfg := config.Read()
	migration(cfg.Database)
	d := db.New(cfg.Database)
	user := store.NewUser(d)

	m := model.User{
		Email:    "parham.alvani@gmail.com",
		Password: "1373",
	}

	assert.Nil(t, user.Insert(m))

	u, err := user.Retrieve(m)

	assert.Nil(t, err)

	assert.Equal(t, m.Email, u.Email)
	assert.Equal(t, m.Password, u.Password)
}

func TestURL(t *testing.T) {
	cfg := config.Read()
	migration(cfg.Database)
	d := db.New(cfg.Database)
	user := store.NewUser(d)

	m := model.User{
		ID:       1,
		Email:    "elahe.dstn@gmail.com",
		Password: "1373",
	}

	if err := user.Insert(m); err != nil {
		fmt.Println(err)
	}

	url := store.NewURL(d)

	u := model.URL{
		UserID: 1,
		URL:    "https://www.google.com",
		Period: 2,
	}

	assert.Nil(t, url.Insert(u))
}

func migration(cfg config.Database) {
	db, err := sql.Open("postgres", cfg.Cstring())
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