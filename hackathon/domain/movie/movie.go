package movie

import (
	"github.com/jupemara/ddd-guys/hackathon/score"
)

type Movie struct {
	id              Id
	title           Title
	publishdPeriods []PublishPeriod
	scores          score.Scores
	length          length
	labels          []Label
}

// TODO: Length

func (m Movie) Length() length {
	return m.length
}
