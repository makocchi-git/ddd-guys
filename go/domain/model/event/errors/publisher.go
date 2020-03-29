package errors

type IPublisher interface {
	Publish(ErrorEvent)
	// このメソッドは削除しちゃって、LocalPublisherをDIする部分(つまりmain.goとかで)
	// でsubscbierを追加してあげて、local, remoteのpublisherどちらでも使えるように
	// Publishだけ使うってのもありかなと！
	// とはいえ、remoteに投げるのが主流なので、そもそも不要かもしれません...
	Subscribe(ISubscriber)
}
