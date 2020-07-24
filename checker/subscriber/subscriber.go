package subscriber

import (
	"checker/config"
	"log"
	"user/model"

	"github.com/nats-io/go-nats"
)

type Subscriber struct {
	Nats     *nats.Conn
	NatsCfg  config.Nats
}

func New(nc *nats.Conn, natsCfg config.Nats) Subscriber {
	return Subscriber{
		Nats:     nc,
		NatsCfg:  natsCfg,
	}
}

func (s *Subscriber) Subscribe() {
	c, err := nats.NewEncodedConn(s.Nats, nats.GOB_ENCODER)
	if err != nil {
		log.Fatal(err)
	}

	defer c.Close()

	ch := make(chan model.URL)

	if _, err := c.QueueSubscribe(s.NatsCfg.Topic, s.NatsCfg.Queue, func(s model.Status) {
		ch <- s
	}); err != nil {
		log.Fatal(err)
	}

	s.worker(ch)
}

func (s *Subscriber) worker(ch chan model.Status) {
	counter := 0

	for st := range ch {
		s.Redis.Insert(st)
		counter++

		if counter == s.RedisCfg.Threshold {
			statuses := s.Redis.Flush()
			for i := 0; i < len(statuses); i++ {
				if err := s.Status.Insert(statuses[i]); err != nil {
					fmt.Println(err)
				}
			}

			counter = 0
		}
	}
}