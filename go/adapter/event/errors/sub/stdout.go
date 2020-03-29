package sub

import (
	"log"

	event "github.com/jupemara/ddd-guys/go/domain/model/event/errors"
)

type StdoutLoggingSubscriber struct{}

func NewStdoutLoggingSubscriber() *StdoutLoggingSubscriber {
	return &StdoutLoggingSubscriber{}
}

func (s *StdoutLoggingSubscriber) Handle(e event.ErrorEvent) {
	log.Println(e)
	return
}
