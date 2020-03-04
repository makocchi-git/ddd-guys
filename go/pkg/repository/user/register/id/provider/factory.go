package provider

import (
	"fmt"

	domain "github.com/makocchi-git/ddd-guys/go/pkg/domain/user"
	"github.com/makocchi-git/ddd-guys/go/pkg/util"
)

// 1 行で書こうと思えば書けるんだけど、こっちのほうが "何が valid なのか" が分かりやすくなるかな
// いいと思います!わかりやすい
var validStringSet = util.NewStringSet([]string{
	"uuid",
	"random",
})

func CreateIdProvider(selector string) (domain.IIdProvider, error) {
	if ok := validStringSet.Has(selector); !ok {
		return nil, fmt.Errorf("given id provider isn't supported: %s", selector)
	}
	if selector == "random" {
		return NewRandomStringIDProvider(32), nil
	}

	// デフォルトだと uuid にしているので、なんとなく uuid をこちらに配置
	// switch-case文使っても全然いいとは思うんですが、if で return するコードとifに入らなければreturnするコードって読みやすいのでifに下感じでした!
	return NewUUIDIDProvider(), nil
}
