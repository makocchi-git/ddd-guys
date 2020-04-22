package user_test

import (
	"testing"

	"github.com/jupemara/ddd-guys/go/domain/model/user"
)

func TestNewUser(t *testing.T) {
	cases := map[string]struct {
		FirstName            string
		LastName             string
		ExpectedToOccurError bool
	}{
		"valid in English":             {"John", "Smith", false},
		"valid in multi-byte language": {"太郎", "山田", false},
		"empty last name":              {"John", "", true},
		"empty first name":             {"", "Smith", true},
		"empty both":                   {"", "", true},
	}
	for k, c := range cases {
		id := user.NewId(k)
		_, err := user.NewUser(id, c.FirstName, c.LastName)
		if c.ExpectedToOccurError != (err != nil) {
			t.Errorf(`case: "%s", error: "%s"`, k, err.Error())
		}
	}
}
