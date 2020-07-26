package status_test

import (
	"saver/config"
	"saver/memory"
	"saver/model"
	"saver/store/status"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
)

func TestRedisStatus(t *testing.T) {
	cfg := config.Read()
	r := memory.New(cfg.Redis)
	redis := status.NewRedisStatus(r)

	f := model.Status{
		ID:         1,
		URLID:      2,
		Clock:      time.Now(),
		StatusCode: 200,
	}

	redis.Insert(f)

	s := model.Status{
		ID:         2,
		URLID:      2,
		Clock:      time.Now(),
		StatusCode: 404,
	}

	redis.Insert(s)

	statuses := redis.Flush()

	assert.Equal(t, f.URLID, statuses[0].URLID)
	assert.Equal(t, f.StatusCode, statuses[0].StatusCode)
	assert.Equal(t, s.URLID, statuses[1].URLID)
	assert.Equal(t, s.StatusCode, statuses[1].StatusCode)
}
