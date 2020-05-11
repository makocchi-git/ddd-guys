package movie

import (
	"github.com/google/uuid"
	domain "github.com/jupemara/ddd-guys/hackthon/domain/movie"
)

type UuidIdProvider struct{}

func NewUuidIdProvider() *UuidIdProvider {
	return &UuidIdProvider{}
}

func (u *UuidIdProvider) NextIdentity() (*domain.Id, error) {
	id := uuid.New()
	return domain.NewId(id.String())
}
