package store

import (
	"errors"
	"log"

	"github.com/httpmon/user/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ErrNotFound  = errors.New("this user doesn't exist in the database")
	ErrWrongPass = errors.New("password is not correct")
)

const PasswordHashLen = 14

type User interface {
	Insert(user model.User) error
	Retrieve(user model.User) (model.User, error)
}

type SQLUser struct {
	DB *gorm.DB
}

func NewUser(d *gorm.DB) SQLUser {
	return SQLUser{DB: d}
}

func (u SQLUser) Insert(user model.User) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), PasswordHashLen)
	if err != nil {
		log.Fatal(err)
	}

	user.Password = string(hashPassword)

	result := u.DB.Create(&user)

	return result.Error
}

func (u SQLUser) Retrieve(user model.User) (model.User, error) {
	var us model.User

	u.DB.Where("email = ?", user.Email).First(&us)

	if us.Email == "" {
		return us, ErrNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(us.Password), []byte(user.Password)); err != nil {
		return us, ErrWrongPass
	}

	return us, nil
}
