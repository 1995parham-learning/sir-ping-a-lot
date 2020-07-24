package main

import (
	"user/config"
	"user/db"
	"user/service"
	"user/store"
)

func main() {
	cfg := config.Read()
	d := db.New(cfg.Database)

	api := service.API {
		User:   store.NewUser(d),
		URL:    store.NewURL(d),
		Config: cfg.JWT,
	}
}
