package idprovider

import (
	duser "github.com/makocchi-git/ddd-guys/go/pkg/domain/user"
)

var valid = string{
	"uuid",
	"random",
}

func CreateIdProvider(selector string) (duser.IIdProvider, error) {
	if err := validSelector(selector); err != nil {
		return nil, err;
	}
	if selector == "uuid" {
		return idprovider.NewUUIDIDProvider(), nil
	}
	if selector == "random" {
		return idprovider.NewRandomStringIDProvider(32), nil
	}
}

func validSelector(selector string) error {
	for _, v := range valid {
		if v == selector {
			return nil
		}
	}
	return nil, fmt.Errorf("given selector isn't supported: %s", v)
}