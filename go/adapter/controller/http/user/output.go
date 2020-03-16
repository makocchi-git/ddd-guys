package user

import (
	"fmt"

	usecase "github.com/jupemara/ddd-guys/go/usecase/user"
)

type IOutputPort interface {
	Print(dto usecase.Dto) string
}

type JsonOutputPort struct {
}

func (o *JsonOutputPort) Print(dto usecase.Dto) string {
	return fmt.Sprintf(`{
id: %s
}`, dto.Id)
}
