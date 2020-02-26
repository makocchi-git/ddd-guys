package duser

type IIdProvider interface {
	NextIdentity() (*Id, error)
}
