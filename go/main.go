package main

import (
	"flag"
	"log"
	"os"

	repository "github.com/makocchi-git/ddd-guys/go/pkg/repository/user/register"

	// package名、 repository/user_register/provider/id, repository/user_register/id/provider とかでもいいかもです
	// golangのpackage名規約的にもhttp, netのように短くわかりやすい名前とのことですし
	// >> ちょっと深く掘りすぎかな・・
	//    repository/user/register/id/provider がいいか、 repository/user/id/provier がいいか、悩むけども
	//    id を provide する機能は register の一部で、それ以外からは使われ無さそうという意味でこっちのほうがいいのかな
	// ここは悩ましいですね。まず、"深く掘りすぎ"ということはないです! どんどん掘っていきましょう!
	// >id を provide する機能は register の一部で、それ以外からは使われ無さそうという意味
	// コレDDDっぽい考え方ですね!
	// ただ技術レイヤなので、もし深すぎるかなーと感じたら(実はpackage, ディレクトリ構造に関しては正解がないので、ドメインや、チーム構成、扱う技術領域に依存することもしばしばです)
	// repository/user/id/uuid.go とかでもアリかなと思います!
	idprovider "github.com/makocchi-git/ddd-guys/go/pkg/repository/user/register/id/provider"
	// パッケージエイリアスは愚直にusecaseのがわかりやすいかもですね
	// >> ここも同じく usecase/user/register にしてみる
	//    domain と同じように複数の usecase が出てきた時に分ければいいので、ここでは usecase で定義
	// +1 です!
	usecase "github.com/makocchi-git/ddd-guys/go/pkg/usecase/user/register"
	// >> まとめると
	//    pkg/domain/user/...
	//    pkg/repository/user/...
	//    pkg/usecase/user/...
	//    という風になり、わかりやすくなったかな？
	//    むしろこういう風に合わせちゃうと、なんかしらの縛りっぽくなっちゃうのならば逆効果だけども・・
	// 見やすくなりましたね!! user/ で切るのって domain/user がuserに関する内容だからという理由なので、usecase/user/register.go, usecase/user/delete.go
	// みたいになってuserに関するapplication serviceというのがわかりやすくなったように思います!
	// >なんかしらの縛り
	// 複数のdomain/配下のモデルを扱うようなケースであったとしても、主体となってるのはどれかというのがあると思うので(例えば、user, credit_cardならuserがおそらく主体)、
	// 主体側に合わせるのがよいかなと思います! ただ実際やっていって usecase/user 配下がどんどん増えてくるとかもありえるのでそうなったら、packageを細かく
	// usecase/user_register (もとに戻ってるけどｗ) みたいに部分的に分けていってあげればいいかと!
	// DDDではよく、開発しているうちにさらにしっくり来る名前とか、ディレクトリ名とかがひらめいたりすること多いので、ビジネス要件の文脈や言葉にマッチしてるかどうかが重要になります!!
)

func main() {

	// flags
	var idp = flag.String("id-provider", "uuid", "an id provider [uuid random]")
	var re = flag.String("repository", "csv", "a backend that stores user data [csv stdout]")
	flag.Parse()

	firstName := "John"
	lastName := "Smith"

	// >> factory パターンで実装しなおし
	//    だけどこれでいいのかどうか・・・
	// いいと思います! スッキリして読みやすくなりました!!
	idpRegister, ierr := idprovider.CreateIdProvider(*idp)
	if ierr != nil {
		log.Fatalf("Failed to create an id provider. raw error message: %v", ierr)
		os.Exit(1)
	}
	repositoryRegister, rerr := repository.CreateUserRepository(*re)
	if rerr != nil {
		log.Fatalf("Failed to create repository register. raw error message: %v", rerr)
		os.Exit(1)
	}

	// usecase
	err := usecase.NewUserRegisterUsecase(
		idpRegister,
		repositoryRegister,
	).Execute(
		firstName,
		lastName,
	)
	if err != nil {
		log.Fatalf("Failed to register user. raw error message: %v", err)
	}
	log.Println("Succeeded to register user")
}
