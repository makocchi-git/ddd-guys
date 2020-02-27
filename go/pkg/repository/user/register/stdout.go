package register

import (
	"fmt"
	"os"

	domain "github.com/makocchi-git/ddd-guys/go/pkg/domain/user"
)

// 出力するだけだから debug 用途
type STDOUTRepository struct{}

func NewSTDOUTRepository() *STDOUTRepository {
	return &STDOUTRepository{}
}

func (r *STDOUTRepository) Store(user *domain.User) error {
	id := user.Id()
	firstName := user.Name().FirstName()
	lastName := user.Name().LastName()
	timeStamp := user.RegisteredAt().String()

	fmt.Fprintf(os.Stdout, "%s,%s,%s,%s\n", timeStamp, id, firstName, lastName)
	return nil
}
