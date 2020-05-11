package score

import (
	"fmt"
)

// Scores

type Scores struct {
	values []IScore
}

func (s Scores) Values() []IScore {
	return s.values
}

func NewScores(s []IScore) (*Scores, error) {
	if len(s) == 0 {
		return &Scores{
			values: s,
		}, nil
	}

	var counter map[int]int
	for _, v := range s {
		i := v.ScoreProvider().Id()
		_, ok := counter[i]
		if !ok {
			// new provider
			counter[i] = 0
		}
		counter[i]++
	}

	// ScoreProvider id must be uniq
	for k, v := range counter {
		if v > 1 {
			return nil, fmt.Errorf("%d is duplicated", k)
		}
	}

	return &Scores{
		values: s,
	}, nil
}
