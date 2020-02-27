package register

import (
	"fmt"

	domain "github.com/makocchi-git/ddd-guys/go/pkg/domain/user"
)

type UserRegister struct {
	idProvider     domain.IIdProvider
	userRepository domain.IUserRepository
}

func NewUserRegisterUsecase(idProvider domain.IIdProvider, userRepository domain.IUserRepository) *UserRegister {
	return &UserRegister{
		idProvider:     idProvider,
		userRepository: userRepository,
	}
}

func (u *UserRegister) Execute(firstName, lastName string) error {
	// generate id
	id, err := u.idProvider.NextIdentity()
	if err != nil {
		return fmt.Errorf("Failed to create new id [%v]", err)
	}
	// create user context
	user, err := domain.NewUser(id, firstName, lastName)
	if err != nil {
		return fmt.Errorf("Some given fields are invalid [%v]", err)
	}
	if err = u.userRepository.Store(user); err != nil {
		return fmt.Errorf("Failed to register user [%v]", err)
	}
	return nil
}
