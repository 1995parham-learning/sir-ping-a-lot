package balancer

import (
	"log"
	"server/config"

	"github.com/nats-io/go-nats"
)

func New(n config.Nats) *nats.EncodedConn {
	nc, err := nats.Connect(n.Host)
	if err != nil {
		log.Fatal(err)
	}

	ec, err := nats.NewEncodedConn(nc, nats.GOB_ENCODER)
	if err != nil {
		log.Fatal(err)
	}

	return ec
}
