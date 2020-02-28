package register

import (
	"fmt"

	domain "github.com/makocchi-git/ddd-guys/go/pkg/domain/user"
	"github.com/makocchi-git/ddd-guys/go/pkg/util"
)

// 1 行で書こうと思えば書けるんだけど、こっちのほうが "何が valid なのか" が分かりやすくなるかな
var validStringSet = util.NewStringSet([]string{
	"csv",
	"stdout",
})

func CreateUserRepository(selector string) (domain.IUserRepository, error) {
	if ok := validStringSet.Has(selector); !ok {
		return nil, fmt.Errorf("given selector isn't supported: %s", selector)
	}
	if selector == "stdout" {
		return NewSTDOUTRepository(), nil
	}

	// デフォルトだと csv にしているので、なんとなく csv をこちらに配置
	return NewCSVRepository(), nil
}
