package subscriber

import (
	"checker/config"
	"checker/model"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/nats-io/go-nats"
)

type Checker struct {
	Nats    *nats.EncodedConn
	NatsCfg config.Nats
}

func New(nc *nats.EncodedConn, natsCfg config.Nats) Checker {
	return Checker{
		Nats:    nc,
		NatsCfg: natsCfg,
	}
}

func (c *Checker) Subscribe() {
	ch := make(chan model.URL)

	if _, err := c.Nats.QueueSubscribe(c.NatsCfg.Topic, c.NatsCfg.Queue, func(s model.URL) {
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
	err := c.Nats.Publish("save", s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("In the checker and publish")
	fmt.Println(s)
}

// Only used for testing.
func (c *Checker) PublishURL(u model.URL) {
	err := c.Nats.Publish(c.NatsCfg.Topic, u)
	if err != nil {
		log.Fatal(err)
	}
}

// Only used for testing.
func (c *Checker) SubscribeStatus(st *model.Status) {
	ch := make(chan model.Status)

	if _, err := c.Nats.QueueSubscribe("save", "test", func(s model.Status) {
		ch <- s
	}); err != nil {
		log.Fatal(err)
	}

	f := <-ch

	st.StatusCode = f.StatusCode
}
