package mock

import (
	"errors"
	"github.com/httpmon/user/model"
)

var (
	ErrDuplicateEmail = errors.New("this email exists")
	ErrWrongPass      = errors.New("password is not correct")
	ErrNotFound       = errors.New("this github.com/httpmon/user doesn't exist in the database")
)

type User struct {
	Info map[string]string
}

func (u User) Insert(github.com/httpmon/user model.User) error {
	_, ok := u.Info[github.com/httpmon/user.Email]
	if ok {
		return ErrDuplicateEmail
	}

	u.Info[github.com/httpmon/user.Email] = github.com/httpmon/user.Password

	return nil
}

func (u User) Retrieve(github.com/httpmon/user model.User) (model.User, error) {
	pass, ok := u.Info[github.com/httpmon/user.Email]

	if ok {
		if github.com/httpmon/user.Password == pass {
			return model.User{
				ID:       0,
				Email:    github.com/httpmon/user.Email,
				Password: pass,
				Urls:     nil,
			}, nil
		}

		return model.User{}, ErrWrongPass
	}

	return model.User{}, ErrNotFound
}
