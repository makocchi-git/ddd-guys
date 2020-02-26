package duser

// ここ、Idだけを保存するわけではないので、IUserRepositoryみたいにUserを名前に含めるとinterface名から何をしてくれるやつか判断しやすくなるかもですね!
type IIdStorer interface {
	Store(*User) error
}
