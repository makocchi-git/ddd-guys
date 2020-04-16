package user

import (
	"errors"

	domain "github.com/jupemara/ddd-guys/go/domain/model/user"
)

type UserFindUsecase struct {
	queryService IUserQueryService
}

func NewUserFindUsecase(queryService IUserQueryService) *UserFindUsecase {
	return &UserFindUsecase{queryService: queryService}
}

// 返す値はDTOでもいいし、そのまま返してもいい
// ただそのままのドメインオブジェクトを返すと、このユースケースを使われる場所でもドメインのメソッドが叩かれてしまう可能性があるので注意
// 下記はじめの段階でメソッドシグネチャ(ここでは `Execute(id string) (*Dto, error)` )のみを書いて、
// コンパイルエラーを直していくような形で型定義をしていくスタイルだとビジネスルールに近い形のコードができあがるはずです
func (u *UserFindUsecase) Execute(id string) (*Dto, error) {
	userId := domain.NewId(id)
	dto, err := u.queryService.FindById(userId)
	if err != nil {
		return nil, errors.New("Couldn't find specified user")
	}
	return &dto, nil
}
