package errors

import (
	"time"

	"github.com/google/uuid"
)

type ErrorEvent struct {
	id         string
	userId     string
	err        error
	occurredAt time.Time
}

func NewErrorEvent(
	userId string,
	err error,
) ErrorEvent {
	id, _ := uuid.NewUUID()
	return ErrorEvent{
		id:         id.String(),
		userId:     userId,
		err:        err,
		occurredAt: time.Now(),
	}
}

func (e *ErrorEvent) Id() string {
	return e.id
}

func (e *ErrorEvent) UserId() string {
	return e.userId
}

func (e *ErrorEvent) Err() error {
	return e.err
}

func (e *ErrorEvent) OccurredAt() time.Time {
	return e.occurredAt
}
