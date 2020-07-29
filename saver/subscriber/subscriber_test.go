package subscriber_test

import (
	"saver/balancer"
	"saver/config"
	"saver/mock"
	"saver/model"
	"saver/subscriber"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
)

func TestSubscriber_Subscribe(t *testing.T) {
	cfg := config.Read()
	n := balancer.New(cfg.Nats)

	r := mock.NewRedis()
	d := mock.Status{}

	s := subscriber.New(n, cfg.Nats, &r, cfg.Redis, &d)

	s.Subscribe()

	st := model.Status{
		ID:         1,
		URLID:      1,
		Clock:      time.Now(),
		StatusCode: 200,
	}

	s.Publish(st)

	assert.Equal(t, r.Memory[st.URLID], st.StatusCode)
}
