package movie

import "time"

type PublishPeriod struct {
	contentsProvider ContentsProvider
	startDate        time.Date
	endDate          time.Date
}
