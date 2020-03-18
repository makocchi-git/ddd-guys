package user

import (
	"fmt"

	presentor "github.com/jupemara/ddd-guys/go/usecase/user"
)

type IOutputPort interface {
	Print(dto presentor.Dto) string
}

type JsonOutputPort struct {
}

func (o *JsonOutputPort) Print(dto presentor.Dto) string {
	return fmt.Sprintf(`{
id: %s
}`, dto.Id)
}
