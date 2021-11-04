package store_test

import (
	"fmt"
	"testing"

	"github.com/httpmon/user/config"
	"github.com/httpmon/user/db"
	"github.com/httpmon/user/model"
	"github.com/httpmon/user/store"

	_ "github.com/golang-migrate/migrate/v4/source/file"
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
