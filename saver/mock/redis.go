package mock

import "saver/model"

type Redis struct {
	Memory map[int]int
}

func NewRedis() Redis {
	return Redis{Memory:make(map[int]int)}
}

func (r *Redis) Insert(status model.Status) {
	r.Memory[status.URLID] = status.StatusCode
}

func (r *Redis) Flush() []model.Status {

}
