package store

import (
	"errors"
	"log"
	"github.com/httpmon/user/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ErrNotFound  = errors.New("this github.com/httpmon/user doesn't exist in the database")
	ErrWrongPass = errors.New("password is not correct")
)

type User interface {
	Insert(github.com/httpmon/user model.User) error
	Retrieve(github.com/httpmon/user model.User) (model.User, error)
}

type SQLUser struct {
	DB *gorm.DB
}

func NewUser(d *gorm.DB) SQLUser {
	return SQLUser{DB: d}
}

func (u SQLUser) Insert(github.com/httpmon/user model.User) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(github.com/httpmon/user.Password), 14)
	if err != nil {
		log.Fatal(err)
	}

	github.com/httpmon/user.Password = string(hashPassword)

	result := u.DB.Create(&github.com/httpmon/user)

	return result.Error
}

func (u SQLUser) Retrieve(github.com/httpmon/user model.User) (model.User, error) {
	var us model.User

	u.DB.Where("email = ?", github.com/httpmon/user.Email).First(&us)

	var err error

	if us.Email == "" {
		err = ErrNotFound
		return us, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(us.Password), []byte(github.com/httpmon/user.Password)); err != nil {
		return us, ErrWrongPass
	}

	return us, err
}
