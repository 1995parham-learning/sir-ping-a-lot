package db

import (
	"log"
	"user/config"

	"github.com/jinzhu/gorm"
	"gorm.io/driver/postgres"
)

func New(config config.Database) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.Cstring()), &gorm.Config{})
	if err != nil {
		log.Printf("can not open connection to database due to the following err\n: %s", err)
	}

	return db
}