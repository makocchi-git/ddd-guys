package errors

type IPublisher interface {
	Publish(ErrorEvent)
}
