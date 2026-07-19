package server

import (
	"fmt"
	"log"
	"server/config"
	"server/model"
	"server/store"
	"time"

	"github.com/nats-io/go-nats"
)

type Server struct {
	URL           store.URL
	Duration      int
	NatsConn      *nats.EncodedConn
	NatsCfg       config.Nats
	TimeThreshold int
}

func New(u store.URL, d int, conn *nats.EncodedConn, cfg config.Nats, th int) Server {
	return Server{
		URL:           u,
		Duration:      d,
		NatsConn:      conn,
		NatsCfg:       cfg,
		TimeThreshold: th,
	}
}

func (s *Server) Run() {
	ticker := time.NewTicker(time.Duration(s.Duration) * time.Minute)
	counter := 0

	for {
		<-ticker.C

		counter++

		if counter == s.TimeThreshold {
			counter = 1
		}

		urls, err := s.URL.GetTable()
		if err != nil {
			log.Fatal(err)
		}

		for _, u := range urls {
			if counter%u.Period != 0 {
				continue
			}

			fmt.Println("In the server and the url is")
			fmt.Println(u)

			s.Publish(u)
		}
	}
}

func (s *Server) Publish(u model.URL) {
	err := s.NatsConn.Publish(s.NatsCfg.Topic, u)
	if err != nil {
		log.Fatal(err)
	}
}

// This function is only used for test.
func (s *Server) Subscribe() model.URL {
	ch := make(chan model.URL)

	if _, err := s.NatsConn.QueueSubscribe(s.NatsCfg.Topic, "test", func(s model.URL) {
		ch <- s
	}); err != nil {
		log.Fatal(err)
	}

	return <-ch
}
