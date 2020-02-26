package duser

import (
	"errors"
)

type User struct {
	id   *Id
	name *Name
}

func NewUser(id *Id, firstName, lastName string) (*User, error) {
	name, err := NewName(firstName, lastName)
	if err != nil {
		return nil, err
	}
	return &User{
		id:   id,
		name: name,
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
