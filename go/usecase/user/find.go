package user

import (
	"errors"

	domain "github.com/jupemara/ddd-guys/go/domain/model/user"
)

type UserFindUsecase struct {
	repository domain.IUserRepository
}

func NewUserFindUsecase(repository domain.IUserRepository) *UserFindUsecase {
	return &UserFindUsecase{repository}
}

// 返す値はDTOでもいいし、そのまま返してもいい
// ただそのままのドメインオブジェクトを返すと、このユースケースを使われる場所でもドメインのメソッドが叩かれてしまう可能性があるので注意
func (u *UserFindUsecase) Execute(id string) (*Dto, error) {
	userId := domain.NewId(id)
	user, err := u.repository.FindById(userId)
	if err != nil {
		return nil, errors.New("Couldn't find specified user")
	}
	return &Dto{
		Id:        user.Id(),
		FirstName: user.Name().FirstName(),
		LastName:  user.Name().LastName(),
	}, nil
}

type Dto struct {
	Id        string
	FirstName string
	LastName  string
}
