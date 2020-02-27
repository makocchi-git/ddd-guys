package register

import (
	"fmt"

	domain "github.com/makocchi-git/ddd-guys/go/pkg/domain/user"
)

var valid = []string{
	"csv",
	"stdout",
}

func CreateUserRepository(selector string) (domain.IUserRepository, error) {
	if err := validSelector(selector); err != nil {
		return nil, err
	}
	if selector == "stdout" {
		return NewSTDOUTRepository(), nil
	}

	// デフォルトだと csv にしているので、なんとなく csv をこちらに配置
	return NewCSVRepository(), nil
}

func validSelector(selector string) error {
	for _, v := range valid {
		if v == selector {
			return nil
		}
	}
	return fmt.Errorf("given selector isn't supported: %s", selector)
}
