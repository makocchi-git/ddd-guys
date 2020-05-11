package score

type IScore interface {
	ScoreProvider() ScoreProvider
	Value() int
}
