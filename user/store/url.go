package store

import (
	"github.com/httpmon/user/model"

	"gorm.io/gorm"
)

type URL interface {
	Insert(url model.URL) error
}

type SQLURL struct {
	DB *gorm.DB
}

func NewURL(d *gorm.DB) SQLURL {
	return SQLURL{DB: d}
}

func (u SQLURL) Insert(url model.URL) error {
	result := u.DB.Create(&url)

	return result.Error
}
