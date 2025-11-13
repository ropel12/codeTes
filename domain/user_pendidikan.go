package domain

import (
	"time"
)

type UserPendidikan struct {
	Name       string
	Status     string
	TimeCreate time.Time
	TimeUpdate time.Time
}
