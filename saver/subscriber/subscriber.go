package subscriber

import (
	"log"
	"saver/config"
	"saver/model"
	"saver/store/status"

	"github.com/nats-io/go-nats"
)

type Subscriber struct {
	nc   *nats.Conn
	natsCfg  config.Nats
	r status.RedisStatus
	redisCfg config.Redis
}

func New(nc *nats.Conn, natsCfg config.Nats, r status.RedisStatus, redisCfg config.Redis) Subscriber {
	return Subscriber{
		nc:       nc,
		natsCfg:  natsCfg,
		r:        r,
		redisCfg: redisCfg,
	}
}

func (s *Subscriber) Subscribe() {
	c, err := nats.NewEncodedConn(s.nc, nats.GOB_ENCODER)
	if err != nil {
		log.Fatal(err)
	}

	defer c.Close()

	ch := make(chan model.Status)

	if _, err := c.QueueSubscribe(s.natsCfg.Topic, s.natsCfg.Queue, func(s model.Status) {
		ch <- s
	}); err != nil {
		log.Fatal(err)
	}

	worker(ch, s.r)
}

func worker(ch chan model.Status, r status.RedisStatus) {
	counter := 0

	for s := range ch {
		r.Insert(s)
		counter++

		if counter ==
	}
}