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
// 下記はじめの段階でメソッドシグネチャ(ここでは `Execute(id string) (*Dto, error)` )のみを書いて、
// コンパイルエラーを直していくような形で型定義をしていくスタイルだとビジネスルールに近い形のコードができあがるはずです
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

// こちらのパターンのように上位のレイヤー(UIやhttp handler)変更してほしくない値だけを
// read onlyにしてガードしてあげるというパターンもアリえます
type ReadOnlyIdDto struct {
	id        string
	FirstName string
	LastName  string
}

func NewDto(id string) ReadOnlyIdDto {
	return ReadOnlyIdDto{
		id: id,
	}
}

func (d ReadOnlyIdDto) Id() string {
	return d.id
}
