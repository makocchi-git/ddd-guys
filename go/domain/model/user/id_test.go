package user_test

import (
	"testing"

	"github.com/jupemara/ddd-guys/go/domain/model/user"
)

func TestNewId(t *testing.T) {
	id, err := user.NewId()
	if err != nil {
		t.Errorf("error message: %v", err)
	}
	if len(id.Value()) <= 0 {
		t.Errorf("something wrong within generating UUID. actual value: %s", id.Value())
	}
}
