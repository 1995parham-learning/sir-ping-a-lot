package main

import (
	"server/config"
	"server/db"
	"server/server"
	"server/store"
)

func main() {
	cfg := config.Read()
	d := db.New(cfg.Database)
	s := server.New(store.NewURL(d),1, , cfg.Nats)
}
