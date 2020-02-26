package idprovider

import (
	"errors"

	"github.com/google/uuid"
	duser "github.com/makocchi-git/ddd-guys/go/pkg/domain/user"
)

type UUIDIDProvider struct{}

func NewUUIDIDProvider() *UUIDIDProvider {
	return &UUIDIDProvider{}
}

func (u *UUIDIDProvider) NextIdentity() (*duser.Id, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, errors.New("assertion error")
	}
	return duser.NewId(id.String()), nil
}
