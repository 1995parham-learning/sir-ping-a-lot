package subscriber_test

import (
	"saver/balancer"
	"saver/config"
	"saver/db"
	"saver/mock"
	"saver/store/status"
	"saver/subscriber"
	"testing"
)

func TestSubscriber_Subscribe(t *testing.T) {
	cfg := config.Read()
	d := db.New(cfg.Database)
	n := balancer.New(cfg.Nats)

	r := mock.NewRedis()

	s := subscriber.New(n, cfg.Nats, &r, cfg.Redis, status.NewSQLStatus(d))

	s.Subscribe()
}
