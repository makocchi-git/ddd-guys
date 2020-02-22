package user

// UserIdを生成するのでdomain/model/user配下に作成する
// 返り値は技術的詳細を抽象化するためにドメイン層のIdを返却する
type IIdProvider interface {
	// Idを提供するメソッド名としてNextIdentityは慣習的によく使われます(iDDD本にも登場しています)
	NextIdentity() (*Id, error)
}
