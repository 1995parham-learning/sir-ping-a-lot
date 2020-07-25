package mock

import "user/model"

type Status struct {
}

func (s *Status) Insert(status model.Status) error {
	return nil
}
