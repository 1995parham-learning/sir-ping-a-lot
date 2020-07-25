package authentication

import (
	"testing"
	"user/config"

	"github.com/stretchr/testify/assert"
)

func TestAuthentication(t *testing.T) {
	cfg := config.Read()
	id := 1

	token, err := CreateToken(id, cfg.JWT)
	assert.Nil(t, err)

	b, d := ValidateToken(token, cfg.JWT)
	assert.True(t, b)
	assert.Equal(t, id, d)
}