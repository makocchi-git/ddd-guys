package pub

import (
	event "github.com/jupemara/ddd-guys/go/domain/model/event/errors"
)

type LocalPublisher struct {
	subscribers []event.ISubscriber
}

// スレッドセーフなsingletonとして実装
// func initを使うのもいいかなと思います
var instance = &LocalPublisher{
	subscribers: []event.ISubscriber{},
}

func NewLocalPublisher() *LocalPublisher {
	return instance
}

func (p *LocalPublisher) Publish(e event.ErrorEvent) {
	for _, v := range p.subscribers {
		v.Handle(e)
	}
	return
}

func (p *LocalPublisher) Subscribe(sub event.ISubscriber) {
	p.subscribers = append(p.subscribers, sub)
	return
}
