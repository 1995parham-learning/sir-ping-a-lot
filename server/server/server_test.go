package server_test

import (
	"server/balancer"
	"server/config"
	"server/mock"
	"server/server"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestServer_Run(t *testing.T) {
	table := map[string]int{"elahe.dstn@gmail.com": 1}
	cfg := config.Read()

	s := server.New(mock.New(table), 1, balancer.New(cfg.Nats), cfg.Nats, cfg.TimeThreshold)

	go s.Run()

	m := s.Subscribe()

	assert.Equal(t, m.URL, "elahe.dstn@gmail.com")
}
