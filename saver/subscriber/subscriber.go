package subscriber

import (
	"log"
	"saver/config"

	"github.com/nats-io/go-nats"
)

func Subscribe(nc *nats.Conn, cfg config.Nats, r status.RedisStatus) {
	c, err := nats.NewEncodedConn(nc, nats.GOB_ENCODER)
	if err != nil {
		log.Fatal(err)
	}

	defer c.Close()

	ch := make(chan model.URL)

	if _, err := c.QueueSubscribe(cfg.Topic, cfg.Queue, func(u model.URL) {
		ch <- u
	}); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 3; i++ {
		go worker(ch, r)
	}

	select {}
}