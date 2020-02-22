package user

type Id struct {
	/*
	   - Id型の構造体を作ることで、技術的な詳細(ここではUUIDを使ってIdを生成していること)を隠蔽する
	   - ビジネス的にシステム一意のユーザオブジェクトを作成したいという要件があったとき、Idをどう生成するかはメインの問題ではない
	   - このUUIDライブラリに万が一致命的なバグがあって違うライブラリに差し替えるときも変更箇所はここだけになる
	   - もし異なる方法でIdを生成する(例えばDBに生成させたり)ケースであっても、Id型を作っておけば差し替えが可能
	*/
	value string
}

func NewId(value string) *Id {
	return &Id{
		value: value,
	}
}

func (id *Id) Value() string {
	return id.value
}
