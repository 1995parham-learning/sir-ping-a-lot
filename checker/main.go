package main

import (
	"checker/balancer"
	"checker/config"
	"checker/subscriber"
)

func main() {
	cfg := config.Read()
	s := subscriber.New(balancer.New(cfg.Nats), cfg.Nats)

	s.Subscribe()
}
