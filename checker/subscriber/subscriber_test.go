package subscriber_test

import (
	"checker/balancer"
	"checker/config"
	"checker/model"
	"checker/subscriber"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
)

func TestChecker_Subscribe(t *testing.T) {
	cfg := config.Read()
	s := subscriber.New(balancer.New(cfg.Nats), cfg.Nats)

	go s.Subscribe()

	st := model.Status{
		ID:         0,
		URLID:      0,
		Clock:      time.Time{},
		StatusCode: 0,
	}

	go s.SubscribeStatus(&st)

	// Let both NATS subscriptions register before publishing; NATS core has no
	// replay, so a message sent before the subscription exists is lost.
	time.Sleep(500 * time.Millisecond)

	s.PublishURL(model.URL{
		ID:       0,
		UserID:   0,
		URL:      "https://www.google.com",
		Period:   0,
		Statuses: nil,
	})

	// Allow the URL fetch (1s timeout) plus the status round-trip to complete.
	time.Sleep(3 * time.Second)

	assert.Equal(t, st.StatusCode, 200)
}
