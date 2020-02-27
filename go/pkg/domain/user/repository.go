package user

// ここ、Idだけを保存するわけではないので、IUserRepositoryみたいにUserを名前に含めるとinterface名から何をしてくれるやつか判断しやすくなるかもですね!
// >> 確かに、いまは Store() しか無いけれども Verify() とか Remove() とか入ってくる可能性はあるね
//    なので Repository にするのが良さそうですね
type IUserRepository interface {
	Store(*User) error
}
