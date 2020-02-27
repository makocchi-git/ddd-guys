package register

import (
	"fmt"
	"os"

	domain "github.com/makocchi-git/ddd-guys/go/pkg/domain/user"
)

// 出力するだけだから debug 用途
type STDOUTBackend struct{}

func NewSTDOUTBackend() *STDOUTBackend {
	return &STDOUTBackend{}
}

func (r *STDOUTBackend) Store(user *domain.User) error {
	id := user.Id()
	firstName := user.Name().FirstName()
	lastName := user.Name().LastName()

	fmt.Fprintf(os.Stdout, "%s,%s,%s\n", id, firstName, lastName)
	return nil
}
