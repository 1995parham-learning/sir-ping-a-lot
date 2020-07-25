package store_test

import (
	"fmt"
	"testing"
	"user/config"
	"user/db"
	"user/model"
	"user/store"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	cfg := config.Read()
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
