package user

import "time"

type RegisteredAt struct {
	timeStamp time.Time
}

func NewRegisteredAt() *RegisteredAt {
	return &RegisteredAt{
		timeStamp: time.Now(),
	}
}

func (r *RegisteredAt) TimeStamp() time.Time {
	return r.timeStamp
}