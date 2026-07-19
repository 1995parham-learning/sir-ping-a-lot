package balancer

import (
	"log"
	"saver/config"

	"github.com/nats-io/go-nats"
)

func New(n config.Nats) *nats.EncodedConn {
	nc, err := nats.Connect(n.Host)
	if err != nil {
		log.Fatal(err)
	}

	c, err := nats.NewEncodedConn(nc, nats.GOB_ENCODER)
	if err != nil {
		log.Fatal(err)
	}

	return c
}
