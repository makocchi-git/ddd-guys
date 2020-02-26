package duser

type Id struct {
	value string
}

func NewId(value string) *Id {
	return &Id{
		value: value,
	}
}

func (id *Id) Value() string {
	return id.value
}
