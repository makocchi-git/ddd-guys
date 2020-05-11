package score

// Scores

func (m Movie) Scores() Scores {
	return m.scores
}

type Scores struct {
	values []Score
}

func (s Scores) Values() {
	return s.values
}

func NewScores(s []IScore) (*Scores, error) {
	if len(s) == 0 {
		return &Scores{
			values: s
		}, nil
	}

	var counter map[int]int
	for _, v := range s {
		i := v.ScoreProvider().Id()
		p, ok := counter[i]
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
		values: s
	}, nil
}

type IScore interface {
	ScoreProvider() ScoreProvider
	Value() int
}

type ScoreProvider struct {
	id int
}

func (s ScoreProvider) Id() int {
	return s.id
}

var ScoreProviderMap [int]string {
	AMAZON: "Amazon Prime Video",
	YAHOO: "Yahoo! Movie",
}

const (
	_ = iota
	AMAZON
	YAHOO
)