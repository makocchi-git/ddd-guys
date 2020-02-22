package user

import (
	"errors"

	"github.com/google/uuid"
	domain "github.com/jupemara/ddd-guys/go/domain/model/user"
)

type UuidRepository struct{}

func NewUuidRepository() *UuidRepository {
	return &UuidRepository{}
}

func (u *UuidRepository) NextIdentity() (*domain.Id, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, errors.New("assertion error")
	}
	return domain.NewId(id.String()), nil
}
