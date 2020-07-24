package subscriber

import "github.com/nats-io/go-nats"

type Subscriber struct {
	Nats     *nats.Conn
	NatsCfg  config.Nats
}