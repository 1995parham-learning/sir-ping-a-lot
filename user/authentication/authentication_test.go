package authentication_test

import (
	"testing"
	"github.com/httpmon/user/authentication"
	"github.com/httpmon/user/config"

	"github.com/stretchr/testify/assert"
)

func TestAuthentication(t *testing.T) {
	cfg := config.Read()
	id := 1

	token, err := authentication.CreateToken(id, cfg.JWT)
	assert.Nil(t, err)

	b, d := authentication.ValidateToken(token, cfg.JWT)
	assert.True(t, b)
	assert.Equal(t, id, d)
}
