package server

import (
	"log"
	"server/store"
	"time"

	"github.com/nats-io/go-nats"
)

type Server struct {
	URL       store.SQLURL
	Duration  int
	NatsConn  *nats.Conn
	NatsCfg   config.Nats
}

func (s *Server) Run() {
	ticker := time.NewTicker(time.Duration(s.Duration) * time.Minute)
	counter := 0

	for {
		<-ticker.C

		urls, err := s.URL.GetTable()
		if err != nil {
			log.Fatal(err)
		}

		for _, u := range urls {
			if counter % u.Period != 0 {
				continue
			}

			s.Publish(u)
		}
	}
}