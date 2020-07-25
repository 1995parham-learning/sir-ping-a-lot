package mock

import (
	"user/model"
)

type URL struct {
	Urls map[string]int
}

func (u URL) Insert(url model.URL) error {
	u.Urls[url.URL] = url.Period

	return nil
}
