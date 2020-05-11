package movie

type Movie struct {
	id              Id
	title           Title
	publishdPeriods []PublishPeriod
	scores          Scores
	length          length
	labels          []Label
}

// ID

func (m Movie) ID() string {
	return m.id.Value()
}

type Id struct {
	value string
}

func (i Id) Value() string {
	return i.value
}

// Title

func (m Movie) Title() string {
	return m.title.Value()
}

type Title struct {
	value string
}

func (t Title) Value() string {
	return t.value
}

// Labels

func (m Movie) Labels() []Label {
	return m.labels
}

type Label struct {
	value string
}

func (l Label) Value() string {
	return l.value
}

// TODO: Length

func (m Movie) Length() length {
	return m.length
}
