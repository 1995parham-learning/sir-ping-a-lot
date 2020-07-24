package subscriber

import (
	"checker/config"
	"checker/model"
	"fmt"
	"log"
	"net/http"
	"time"

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

	if _, err := c.QueueSubscribe(s.NatsCfg.Topic, s.NatsCfg.Queue, func(s model.URL) {
		ch <- s
	}); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 3; i++ {
		go s.worker(ch)
	}

	select {}
}

func (s *Subscriber) worker(ch chan model.URL) {
	counter := 0

	for u := range ch {
		resp, err := http.Get(u.URL)
		if err != nil {
			fmt.Println(err)
		}

		var st model.Status
		st.URLID = u.ID
		st.Clock = time.Now()
		st.StatusCode = resp.StatusCode

		r.Insert(st)
	}
}