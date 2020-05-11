package movie

import (
	"errors"
)

type Id struct {
	value string
}

func (id Id) Value() string {
	return id.value
}

func NewId(value string) (*Id, error) {
	for _, unsatisfied := range []bool{
		len(value) < 6,
		len(value) > 32,
	} {
		if unsatisfied {
			return nil, errors.New("doesn't meet the criteria")
		}
	}
	return &Id{value: value}, nil
}
