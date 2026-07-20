package store_test

import (
	"testing"

	"github.com/httpmon/user/config"
	"github.com/httpmon/user/db"
	"github.com/httpmon/user/model"
	"github.com/httpmon/user/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUser(t *testing.T) {
	t.Parallel()

	cfg := config.Read()
	d := db.New(cfg.Database)
	user := store.NewUser(d)

	//nolint: exhaustivestruct
	m := model.User{
		Email:    "parham.alvani@gmail.com",
		Password: "1373",
	}

	require.NoError(t, user.Insert(m))

	u, err := user.Retrieve(m)
	require.NoError(t, err)

	assert.Equal(t, m.Email, u.Email)
}

func TestURL(t *testing.T) {
	t.Parallel()

	cfg := config.Read()
	d := db.New(cfg.Database)
	user := store.NewUser(d)

	//nolint: exhaustivestruct
	m := model.User{
		ID:       2,
		Email:    "elahe.dstn@gmail.com",
		Password: "1373",
	}

	require.NoError(t, user.Insert(m))

	url := store.NewURL(d)

	//nolint: exhaustivestruct
	u := model.URL{
		UserID: 2,
		URL:    "https://www.google.com",
		Period: 2,
	}

	require.NoError(t, url.Insert(u))
}
