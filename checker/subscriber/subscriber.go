package subscriber

import (
	"checker/balancer"
	"checker/config"
	"checker/model"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/nats-io/nats.go"
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
	ch := make(chan model.URL)

	if _, err := c.Nats.QueueSubscribe(c.NatsCfg.Topic, c.NatsCfg.Queue, func(msg *nats.Msg) {
		var u model.URL
		if err := balancer.Decode(msg.Data, &u); err != nil {
			log.Fatal(err)
		}

		ch <- u
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
		resp, err := fetch(u)

		var st model.Status
		st.URLID = u.ID
		st.Clock = time.Now()

		if err != nil {
			st.StatusCode = http.StatusRequestTimeout
		} else {
			st.StatusCode = resp.StatusCode
		}

		fmt.Println("In the checker the url is")
		fmt.Println(u.URL)

		c.Publish(st)
	}
}

func fetch(u model.URL) (*http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.URL, nil)
	if err != nil {
		fmt.Println(err)
	}

	client := http.DefaultClient

	return client.Do(req)
}

func (c *Checker) Publish(s model.Status) {
	data, err := balancer.Encode(s)
	if err != nil {
		log.Fatal(err)
	}

	if err := c.Nats.Publish("save", data); err != nil {
		log.Fatal(err)
	}

	fmt.Println("In the checker and publish")
	fmt.Println(s)
}

// Only used for testing.
func (c *Checker) PublishURL(u model.URL) {
	data, err := balancer.Encode(u)
	if err != nil {
		log.Fatal(err)
	}

	if err := c.Nats.Publish(c.NatsCfg.Topic, data); err != nil {
		log.Fatal(err)
	}
}

// Only used for testing.
func (c *Checker) SubscribeStatus(st *model.Status) {
	ch := make(chan model.Status)

	if _, err := c.Nats.QueueSubscribe("save", "test", func(msg *nats.Msg) {
		var s model.Status
		if err := balancer.Decode(msg.Data, &s); err != nil {
			log.Fatal(err)
		}

		ch <- s
	}); err != nil {
		log.Fatal(err)
	}

	f := <-ch

	st.StatusCode = f.StatusCode
}
