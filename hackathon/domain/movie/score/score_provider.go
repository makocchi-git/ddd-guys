package score

type ScoreProvider struct {
	id int
}

func (s ScoreProvider) Id() int {
	return s.id
}

var ScoreProviderMap = map[int]string{
	AMAZON: "Amazon Prime Video",
	YAHOO:  "Yahoo! Movie",
}

const (
	_ = iota
	AMAZON
	YAHOO
)
