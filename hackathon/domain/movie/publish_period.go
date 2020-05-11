package movie

import "time"

type PublishPeriod struct {
	contentsProvider ContentsProvider
	startDate        time.Time
	endDate          time.Time
}
