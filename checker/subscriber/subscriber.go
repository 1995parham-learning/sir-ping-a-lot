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

type Checker struct {
	Nats    *nats.Conn
	NatsCfg config.Nats
}

func New(nc *nats.Conn, natsCfg config.Nats) Checker {
	return Checker{
		Nats:    nc,
		NatsCfg: natsCfg,
	}
}

func (c *Checker) Subscribe() {
	ec, err := nats.NewEncodedConn(c.Nats, nats.GOB_ENCODER)
	if err != nil {
		log.Fatal(err)
	}

	defer ec.Close()

	ch := make(chan model.URL)

	if _, err := ec.QueueSubscribe(c.NatsCfg.Topic, c.NatsCfg.Queue, func(s model.URL) {
		ch <- s
	}); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 3; i++ {
		go c.worker(ch)
	}

	select {}
}

//nolint: bodyclose
func (c *Checker) worker(ch chan model.URL) {
	for u := range ch {
		resp, err := http.Get(u.URL)
		if err != nil {
			fmt.Println(err)
		}

		var st model.Status
		st.URLID = u.ID
		st.Clock = time.Now()
		st.StatusCode = resp.StatusCode

		fmt.Println("In the checker the url is")
		fmt.Println(u.URL)

		c.Publish(st)
	}
}

func (c *Checker) Publish(s model.Status) {
	ec, err := nats.NewEncodedConn(c.Nats, nats.GOB_ENCODER)
	if err != nil {
		log.Fatal(err)
	}

	err = ec.Publish("save", s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("In the checker and publish")
	fmt.Println(s)
}
