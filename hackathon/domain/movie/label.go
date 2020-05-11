package movie

func (m Movie) Labels() []Label {
	return m.labels
}

type Label struct {
	value string
}

func (l Label) Value() string {
	return l.value
}
