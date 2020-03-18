package user

import (
	"errors"

	domain "github.com/jupemara/ddd-guys/go/domain/model/user"
)

type UserRegisterUsecase struct {
	// 技術的な実装(具象)に依存するのではなく、ビジネス的成約、契約(ここでは、ユーザを保存するというインターフェイス)にのみ依存する
	// とにかく技術的な詳細はクラスや外のモジュールに押し込んでいくというのがポイントですね
	idProvider     domain.IIdProvider
	UserRepository domain.IUserRepository
}

func NewUserRegisterUsecase(
	idProvider domain.IIdProvider,
	repository domain.IUserRepository,
) *UserRegisterUsecase {
	return &UserRegisterUsecase{
		idProvider:     idProvider,
		UserRepository: repository,
	}
}

// アプリケーションサービスを実行する際の引数はプリミティブ型のみ(もしくは後述のCommandオブジェクト)を使う
func (u *UserRegisterUsecase) Execute(firstName, lastName string) error {
	id, err := u.idProvider.NextIdentity()
	if err != nil {
		return errors.New("Failed to create new id")
	}
	// ここでまずビジネス条件を違反するインスタンスはそもそも作らせない
	user, err := domain.NewUser(id, firstName, lastName)
	if err != nil {
		return errors.New("Some given fields are invalid")
	}
	if err = u.UserRepository.Store(user); err != nil {
		return errors.New("Failed to register user")
	}
	return nil
}

// Commandオブジェクトを使ったDTOの実装例
// CommandオブジェクトとDTOの定義は結構曖昧ですが、doc/3.mdに定義を書いておきました
// プレーンオブジェクトとして受け取るのを意図としているのであえて値渡しをしています(引数に与えることで副作用がないことを保証する)。ただパフォーマンスに影響があるのであれば参照渡しでgetterのみバージョンを使いましょう
func (u *UserRegisterUsecase) ExecuteWithCommand(command Command) error {
	id, err := u.idProvider.NextIdentity()
	if err != nil {
		return errors.New("Failed to create new id")
	}
	// ここでまずビジネス条件を違反するインスタンスはそもそも作らせない
	user, err := domain.NewUser(id, command.FirstName, command.LastName)
	if err != nil {
		return errors.New("Some given fields are invalid")
	}
	if err = u.UserRepository.Store(user); err != nil {
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
type Command struct {
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
