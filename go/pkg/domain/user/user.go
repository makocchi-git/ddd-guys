package user

import (
	"errors"
	"time"
)

type User struct {
	id           *Id
	name         *Name
	registeredAt *RegisteredAt
}

func NewUser(id *Id, firstName, lastName string) (*User, error) {
	name, err := NewName(firstName, lastName)
	if err != nil {
		return nil, err
	}
	return &User{
		id:           id,
		name:         name,
		registeredAt: NewRegisteredAt(),
	}, nil
}

func (u *User) ChangeName(firstName, lastName string) error {
	name, err := NewName(firstName, lastName)
	if err != nil {
		return errors.New("assertion error")
	}
	u.name = name
	return nil
}

func (u *User) Name() *Name {
	return u.name
}

func (u *User) Id() string {
	return u.id.Value()
}

func (u *User) RegisteredAt() time.Time {
	return u.registeredAt.TimeStamp()
}
