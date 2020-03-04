package user

import (
	"fmt"

	domain "github.com/jupemara/ddd-guys/go/domain/model/user"
)

type UserUpdateUsecase struct {
	repository domain.IUserRepository
}

func NewUserUpdateUsecase(repository domain.IUserRepository) *UserUpdateUsecase {
	return &UserUpdateUsecase{repository}
}

func (u *UserUpdateUsecase) Execute(command UpdateCommand) error {
	/*
		ユースケースにユーザのファーストネームは3文字以下であるのようなルールを記載せずに、
		こういったルールは値オブジェクトのコンストラクタに定義して、ガードする
		if len(command.FirstName()) < 3 {
			return errors.New("validation error")
		}
	*/
	id := domain.NewId(command.Id())
	// UIなどから詰められた値を利用するのではなく、まず変更前のオブジェクトを取得し、
	// その後ドメインオブジェクトのメソッド(今回ならUser.ChangeName)を使って変更を加えていく
	user, err := u.repository.FindById(id)
	if err != nil {
		return err // TODO: handle error
	}
	if len(command.FirstName()) > 0 && len(command.LastName()) == 0 {
		// 変更を行うのはrepositoryではなくドメインオブジェクトのメソッド
		// repositoryはあくまで永続化だけを行う
		err := user.ChangeName(
			command.FirstName(),
			user.Name().LastName(),
		)
		if err != nil {
			return err // TODO: handle error
		}
	}
	if len(command.LastName()) > 0 && len(command.FirstName()) == 0 {
		err := user.ChangeName(
			user.Name().FirstName(),
			command.LastName(),
		)
		if err != nil {
			return err // TODO: handle error
		}
	}
	if len(command.FirstName()) > 0 && len(command.LastName()) > 0 {
		err := user.ChangeName(
			command.FirstName(),
			command.LastName(),
		)
		if err != nil {
			return err // TODO: handle error
		}
	}
	// ここは不用意にDBへのアクセスを防いでいるだけなので、実装する必要はないかもしれません
	if len(command.FirstName()) == 0 && len(command.LastName()) == 0 {
		return nil
	}
	err = u.repository.Update(user)
	if err != nil {
		return fmt.Errorf("Failed to update [%v]", err)
	}
	return nil
}

func NewCommand(id, firstName, lastName string) UpdateCommand {
	return UpdateCommand{id, firstName, lastName}
}

// commandオブジェクトの場合は、DTOと違って、usecase内で値を変更できる可能性を排除したいので、
// すべてread onlyなフィールドにしてガードします
type UpdateCommand struct {
	id        string
	firstName string
	lastName  string
}

func (c *UpdateCommand) Id() string {
	return c.id
}

func (c *UpdateCommand) FirstName() string {
	return c.firstName
}

func (c *UpdateCommand) LastName() string {
	return c.lastName
}
