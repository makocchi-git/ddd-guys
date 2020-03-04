package register

import (
	"fmt"

	domain "github.com/makocchi-git/ddd-guys/go/pkg/domain/user"
	service "github.com/makocchi-git/ddd-guys/go/pkg/domain/service/user"
)

type UserRegister struct {
	idProvider     domain.IIdProvider
	userRepository domain.IUserRepository
	userService    service.UserExistsService
}

func NewUserRegisterUsecase(
	idProvider		domain.IIdProvider,
	userRepository 	domain.IUserRepository,
	userService 	service.UserExistsService,
) *UserRegister {
	return &UserRegister{
		idProvider:     idProvider,
		userRepository: userRepository,
		userService: 	userService,
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
	// ここにStoreする前に重複排除コードを入れる
	if !u.userService.Exists(user.Name()) {
		return fmt.Errorf("prohibited to register deplicated name user: %s", user.Name().FullName()) // FirstNameとLastNameを分けて出力してもいいかも
	}
	if err = u.userRepository.Store(user); err != nil {
		return fmt.Errorf("Failed to register user [%v]", err)
	}
	return nil
}
