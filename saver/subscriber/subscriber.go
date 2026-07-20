package subscriber

import (
	"fmt"
	"log"
	"saver/balancer"
	"saver/config"
	"saver/model"
	"saver/store/status"

	"github.com/nats-io/nats.go"
)

type Subscriber struct {
	Nats     *nats.Conn
	NatsCfg  config.Nats
	Redis    status.Memory
	RedisCfg config.Redis
	Status   status.Status
}

func New(nc *nats.Conn, natsCfg config.Nats, r status.Memory,
	redisCfg config.Redis, s status.Status) Subscriber {
	return Subscriber{
		Nats:     nc,
		NatsCfg:  natsCfg,
		Redis:    r,
		RedisCfg: redisCfg,
		Status:   s,
	}
}

func (s *Subscriber) Subscribe() {
	ch := make(chan model.Status)

	if _, err := s.Nats.QueueSubscribe(s.NatsCfg.Topic, s.NatsCfg.Queue, func(msg *nats.Msg) {
		var st model.Status
		if err := balancer.Decode(msg.Data, &st); err != nil {
			log.Fatal(err)
		}

		ch <- st
	}); err != nil {
		log.Fatal(err)
	}

	s.worker(ch)
}

func (s *Subscriber) Publish(st model.Status) {
	data, err := balancer.Encode(st)
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Nats.Publish(s.NatsCfg.Topic, data); err != nil {
		log.Fatal(err)
	}

	fmt.Println("In the checker and publish")
	fmt.Println(st)
}

func (s *Subscriber) worker(ch chan model.Status) {
	counter := 0

	for st := range ch {
		fmt.Println("In the saver")
		fmt.Println(st)
		s.Redis.Insert(st)
		counter++

		fmt.Println(counter)

		if counter == s.RedisCfg.Threshold {
			statuses := s.Redis.Flush()
			for i := 0; i < len(statuses); i++ {
				if err := s.Status.Insert(statuses[i]); err != nil {
					fmt.Println(err)
				}
			}

			fmt.Println("flush")

			counter = 0
		}
	}
}
