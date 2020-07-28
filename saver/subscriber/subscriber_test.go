package subscriber_test

import (
	"saver/balancer"
	"saver/config"
	"saver/db"
	"saver/memory"
	"saver/subscriber"
	"testing"
)

func TestSubscriber_Subscribe(t *testing.T) {
	cfg := config.Read()
	d := db.New(cfg.Database)
	r := memory.New(cfg.Redis)
	n := balancer.New(cfg.Nats)

	subscriber.New(n, cfg.Nats, )
}
