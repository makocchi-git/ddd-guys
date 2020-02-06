package user_register

import (
	"errors"

	domain "github.com/jupemara/ddd-guys/go/domain/model/user"
)

type UserRegister struct {
	// 技術的な実装(具象)に依存するのではなく、ビジネス的成約、契約(ここでは、ユーザを保存するというインターフェイス)にのみ依存する
	// とにかく技術的な詳細はクラスや外のモジュールに押し込んでいくというのがポイントですね
	UserRepository domain.IUserRepository
}

func NewUserRegisterUsecase(repository domain.IUserRepository) *UserRegister {
	return &UserRegister{
		UserRepository: repository,
	}
}

// アプリケーションサービスを実行する際の引数はプリミティブ型のみを使う
func (u *UserRegister) Execute(firstName, lastName string) error {
	// ここでまずビジネス条件を違反するインスタンスはそもそも作らせない
	user, err := domain.NewUser(firstName, lastName)
	if err != nil {
		return errors.New("Some given fields are invalid")
	}
	err = u.UserRepository.Store(user)
	if err != nil {
		return errors.New("Failed to register user")
	}
	return nil
}

// Commandオブジェクトを使ったDTOの実装例
// CommandオブジェクトとDTOの定義は結構曖昧ですが、doc/3.mdに定義を書いておきました
// プレーンオブジェクトとして受け取るのを意図としているのであえて値渡しをしています(引数に与えることで副作用がないことを保証する)。ただパフォーマンスに影響があるのであれば参照渡しでgetterのみバージョンを使いましょう
func (u *UserRegister) ExecuteWithCommand(command command) error {
	// ここでまずビジネス条件を違反するインスタンスはそもそも作らせない
	user, err := domain.NewUser(command.FirstName, command.LastName)
	if err != nil {
		return errors.New("Some given fields are invalid")
	}
	err = u.UserRepository.Store(user)
	if err != nil {
		return errors.New("Failed to register user")
	}
	return nil
}

// getterのみでいいケースも多いので、コンストラクタとgetterの実装も結構あります
// ただ結構書いてるとめんどくなってくるので、受け渡しだけ限定なので(ビジネスロジックが入らないオブジェクトなので)、すべてpublic属性にしてしまうというのもありだと思います(golangなら自分はそうするかも。)
/*
ただ例えばTypeScriptだと
class Command {
  constructor(
    public readonly firstName: string,
    public readonly lastName: string,
  ) {}
}
こう書くだけで、
const command = new Command('John', 'Smith');
command.firstName // 'John'をget
command.firstName = 'Johhhhhhhhhhn' // コンパイルエラー
こんな感じにできるので、ワイならTypeScriptでこの形で書くと思います
*/
type command struct {
	FirstName string
	LastName  string
}

// Getterのみversion, ただ見てもらえると分かる通り、記述量が増えてめんどい
type commandOnlyGetter struct {
	firstName string
	lastName  string
}

func (c *commandOnlyGetter) FirstName() string {
	return c.firstName
}
func (c *commandOnlyGetter) LastName() string {
	return c.lastName
}
