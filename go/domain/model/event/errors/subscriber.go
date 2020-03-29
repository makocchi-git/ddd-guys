package errors

type ISubscriber interface {
	Handle(ErrorEvent)
}
