package user

// interfaceはDIPを実現させるため、またよりビジネスのコアに依存させるために円の内側(つまりドメイン層)に定義します
type IUserRepository interface {
	Store(*User) error
	FindById(*Id) (*User, error)
	// ChangeUserNameとかUpdateFirstNameみたいにひとつひとつのフィールドごとに
	// application serviceやrepositoryのメソッドを増やすのはアンチパターン
	// 受け取るのはusecaseですでに変更済みのUserオブジェクトだけを受け取る
	Update(*User) error
}
