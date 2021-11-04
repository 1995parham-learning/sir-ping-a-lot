package authentication

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/httpmon/user/config"
)

func CreateToken(id int, cfg config.JWT) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["github.com/httpmon/user_id"] = id
	atClaims["exp"] = time.Now().Add(time.Minute * time.Duration(cfg.Expiration)).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(cfg.Secret))
	if err != nil {
		return "", fmt.Errorf("signing token failed %w", err)
	}

	return token, nil
}

func ValidateToken(token string, cfg config.JWT) (in bool, i int) {
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.Secret), nil
	})
	if err != nil {
		return false, 0
	}

	auth, ok := claims["authorized"].(bool)
	if !ok {
		return false, 0
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return false, 0
	}

	id, ok := claims["github.com/httpmon/user_id"].(float64)
	if !ok {
		return false, 0
	}

	if auth && exp > float64(time.Now().Unix()) {
		return true, int(id)
	}

	return false, int(id)
}
