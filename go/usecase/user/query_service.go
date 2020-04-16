package user

import (
	domain "github.com/jupemara/ddd-guys/go/domain/model/user"
)

type IUserQueryService interface {
	FindById(*domain.Id) (Dto, error)
}

type Dto struct {
	Id        string
	FirstName string
	LastName  string
	FullName  string
}

// こちらのパターンのように上位のレイヤー(UIやhttp handler)変更してほしくない値だけを
// read onlyにしてガードしてあげるというパターンもアリえます
type ReadOnlyIdDto struct {
	id        string
	FirstName string
	LastName  string
	FullName  string
}

func NewDto(id string) ReadOnlyIdDto {
	return ReadOnlyIdDto{
		id: id,
	}
}

func (d ReadOnlyIdDto) Id() string {
	return d.id
}
