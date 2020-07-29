package mock

import "server/model"

type URL struct {
	Table map[string]int
}

func New(t map[string]int) URL {
	return URL{Table: t}
}

func (u URL) GetTable() ([]model.URL, error) {
	models := make([]model.URL, 0)

	for k, v := range u.Table {
		models = append(models, model.URL{
			ID:       0,
			UserID:   0,
			URL:      k,
			Period:   v,
			Statuses: nil,
		})
	}

	return models, nil
}
