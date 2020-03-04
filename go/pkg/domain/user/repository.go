package user

// ここ、Idだけを保存するわけではないので、IUserRepositoryみたいにUserを名前に含めるとinterface名から何をしてくれるやつか判断しやすくなるかもですね!
// >> 確かに、いまは Store() しか無いけれども Verify() とか Remove() とか入ってくる可能性はあるね
//    なので Repository にするのが良さそうですね

// domain/service/user/exists.go でも触れましたが、
// チャットでもらってた質問、
// - そもそも重複チェックみたいなことをやりたい場合にロジックは repository の Store() 内でいいのか？
// - Store() とは別に Check() みたいに分けたほうがいいのか？
// について、後者が正解です!! なぜならrepositoryは、永続層の機能の提供であり、domain層のinterfaceを目指すべき存在であり
// domain層のinterfaceとはつまり、ビジネス要件だからです!!
// "やりたい場合にロジックは"ある通り、ここでいうロジックはビジネスロジックということになるかとオモワれます

// この場合の実装パターンとしては2つで
// 1. FindByXXXメソッドを提供して、repository層ではあくまでユーザを見つけてくる部分だけを実装する
// 2. Existsメソッドを提供して、domain serviceでwrapして、application serviceで使用する
// 1はメソッドの使い勝手がよくなるという反面、いろんなところで使ってしまいがちで、変更の柔軟性が失われる可能性があるというデメリットがあります
// 2は逆にメソッドの使用が比較的局所的になる反面、使い回しができず、
// 結局ほぼ同じようなコードをrepositoryそうでFindByXXX, Existsというメソッドで実装してしまうきらいがあります
type IUserRepository interface {
	Store(*User) error
	// パターン1 今回は名前で重複チェックを行うということですので
	FindByName(Name) (User, error)
	// パターン2 ここの引数にはUserを指定してもよいのですが, おそらくビジネスルールで
	// - 氏名による重複チェックを行う、とか同じ氏名は登録ができない
	// - メールアドレスによる重複チェックを行う
	// のようにどのユーザの属性で重複チェックを行うことがビジネスルールで決められると思うので、それに従って引数を指定します
	// またこの時、premitive型は極端な話ですが、repository.Exists(user.UserId().Value())を呼び出してもコンパイルが通ってしまうので、
	// ドメイン上のオブジェクト(ここではName値オブジェクト)を引数に指定するのがベターです
	Exists(Name) bool
}
