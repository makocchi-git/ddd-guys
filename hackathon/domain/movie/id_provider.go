package movie

type IIdProvider interface {
	NextIdentity() (*Id, error)
}
