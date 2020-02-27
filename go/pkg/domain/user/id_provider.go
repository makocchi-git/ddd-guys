package user

type IIdProvider interface {
	NextIdentity() (*Id, error)
}
