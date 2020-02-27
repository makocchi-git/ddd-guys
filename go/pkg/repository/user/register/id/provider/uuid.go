package provider

import (
	"errors"

	"github.com/google/uuid"

	domain "github.com/makocchi-git/ddd-guys/go/pkg/domain/user"
)

type UUIDIDProvider struct{}

func NewUUIDIDProvider() *UUIDIDProvider {
	return &UUIDIDProvider{}
}

func (u *UUIDIDProvider) NextIdentity() (*domain.Id, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, errors.New("assertion error")
	}
	return domain.NewId(id.String()), nil
}
