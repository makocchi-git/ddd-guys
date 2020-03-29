package errors

type IPublisher interface {
	Publish(ErrorEvent)
	Subscribe(ISubscriber)
}
