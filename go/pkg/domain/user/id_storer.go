package duser

type IIdStorer interface {
	Store(*User) error
}
