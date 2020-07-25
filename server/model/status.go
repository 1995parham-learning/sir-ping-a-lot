package model

import "time"

type Status struct {
	ID         int
	URLID      int
	Clock      time.Time
	StatusCode int
}
