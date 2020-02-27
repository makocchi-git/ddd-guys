package register

import (
	"encoding/csv"
	"os"

	domain "github.com/makocchi-git/ddd-guys/go/pkg/domain/user"
)

type CSVBackend struct{}

func NewCSVBackend() *CSVBackend {
	return &CSVBackend{}
}
func (r *CSVBackend) Store(user *domain.User) error {
	id := user.Id()
	firstName := user.Name().FirstName()
	lastName := user.Name().LastName()
	file, err := os.OpenFile("/tmp/users.csv", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	defer file.Close()
	if err != nil {
		return err
	}
	writer := csv.NewWriter(file)
	err = writer.Write([]string{id, firstName, lastName})
	if err != nil {
		return err
	}
	writer.Flush()
	return nil
}
