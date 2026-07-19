package main

import (
	"server/balancer"
	"server/config"
	"server/db"
	"server/server"
	"server/store"
)

func main() {
	cfg := config.Read()
	d := db.New(cfg.Database)
	s := server.New(store.NewURL(d), 1, balancer.New(cfg.Nats), cfg.Nats, cfg.TimeThreshold)

	s.Run()
}
