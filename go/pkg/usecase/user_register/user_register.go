package uregister

import (
	"fmt"

	duser "github.com/makocchi-git/ddd-guys/go/pkg/domain/user"
)

type UserRegister struct {
	// idProvider : generate id
	// idStoreer : store id
	idProvider duser.IIdProvider
	idStorer   duser.IIdStorer
}

func NewUserRegisterUsecase(idProvider duser.IIdProvider, idStorer duser.IIdStorer) *UserRegister {
	return &UserRegister{
		idProvider: idProvider,
		idStorer:   idStorer,
	}
}

func (u *UserRegister) Execute(firstName, lastName string) error {
	// generate id
	id, err := u.idProvider.NextIdentity()
	if err != nil {
		return fmt.Errorf("Failed to create new id [%v]", err)
	}
	// create user context
	user, err := duser.NewUser(id, firstName, lastName)
	if err != nil {
		return fmt.Errorf("Some given fields are invalid [%v]", err)
	}
	if err = u.idStorer.Store(user); err != nil {
		return fmt.Errorf("Failed to register user [%v]", err)
	}
	return nil
}
