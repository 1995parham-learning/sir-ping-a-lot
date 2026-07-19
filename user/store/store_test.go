package store_test

import (
	"testing"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/httpmon/user/config"
	"github.com/httpmon/user/db"
	"github.com/httpmon/user/model"
	"github.com/httpmon/user/store"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	t.Parallel()

	cfg := config.Read()
	d := db.New(cfg.Database)
	user := store.NewUser(d)

	// nolint: exhaustivestruct
	m := model.User{
		Email:    "parham.alvani@gmail.com",
		Password: "1373",
	}

	assert.NoError(t, user.Insert(m))

	u, err := user.Retrieve(m)
	assert.NoError(t, err)

	assert.Equal(t, m.Email, u.Email)
}

func TestURL(t *testing.T) {
	t.Parallel()

	cfg := config.Read()
	d := db.New(cfg.Database)
	user := store.NewUser(d)

	// nolint: exhaustivestruct
	m := model.User{
		ID:       2,
		Email:    "elahe.dstn@gmail.com",
		Password: "1373",
	}

	assert.NoError(t, user.Insert(m))

	url := store.NewURL(d)

	// nolint: exhaustivestruct
	u := model.URL{
		UserID: 2,
		URL:    "https://www.google.com",
		Period: 2,
	}

	assert.NoError(t, url.Insert(u))
}
