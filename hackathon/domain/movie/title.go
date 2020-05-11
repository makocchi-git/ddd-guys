package movie

func (m Movie) Title() string {
	return m.title.Value()
}

type Title struct {
	value string
}

func (t Title) Value() string {
	return t.value
}
