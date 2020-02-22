package user

import (
	"errors"

	"github.com/google/uuid"
	domain "github.com/jupemara/ddd-guys/go/domain/model/user"
)

// golangには`implements`的なやつがないので、どのinterfaceになるのかがわかりやすいように接尾的にinterfaceの名前をつけました
type UuidIdProvider struct{}

func NewUuidIdProvider() *UuidIdProvider {
	return &UuidIdProvider{}
}

func (u *UuidIdProvider) NextIdentity() (*domain.Id, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, errors.New("assertion error")
	}
	return domain.NewId(id.String()), nil
}
