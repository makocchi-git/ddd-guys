package main

import (
	"flag"
	"log"
	"os"

	// application serviceがuserだけにとどまっているなら、パッケージエイリアスは僕は"domain"って使ったりしますね
	duser "github.com/makocchi-git/ddd-guys/go/pkg/domain/user"

	backend "github.com/makocchi-git/ddd-guys/go/pkg/repository/user_register/backend"
	// package名、 repository/user_register/provider/id, repository/user_register/id/provider とかでもいいかもです
	// golangのpackage名規約的にもhttp, netのように短くわかりやすい名前とのことですし
	idprovider "github.com/makocchi-git/ddd-guys/go/pkg/repository/user_register/id_provider"

	// パッケージエイリアスは愚直にusecaseのがわかりやすいかもですね
	register "github.com/makocchi-git/ddd-guys/go/pkg/usecase/user_register"
)

func main() {
	// go っぽく pkg/ にいろいろ入れてみた
	// pkg/ 確かにgolangぽくていいですね!
	// 個人的にはこういう構成の方が見やすいかも
	// でも package 名は苦労しますな・・
	// 今回は domain/user にして domainuser という意味で "duser" にしてある
	// usecase の package 名は usecase の u を接頭語としてみたけども・・
	// 自分のコードがpackage user_registerになってるので恐縮ですけど、usecase/user/register.go とかにしてdomainレベルと同じ粒度で切るのもありかなと思います!

	// 試しに Random String で生成する IDP と保存するんじゃなくて STDOUT に出力する Store() を作ってみた
	// こういうドメインに微妙に関係ないところの実装ってどこに置くか悩ましいんですが、debug入れたいというのを要件のひとつに入れるとするとloggerを実装して、usecaseのひとつとしてloggerにdebugを履かせるのもありかなと!!
	// repository の directory 構成的に .../backend と .../id_provier みたいに細かくしたんだけど
	// そういうのはアリなのかな？
	// repositoryはあくまでビジネスユースを満たすためのapplication serviceで使われる実態として定義されるという前提に基づくと
	// "backend"はやや抽象的すぎる感じがしていて、repository/user_register/db/csv.go とかもっと愚直でわかりやすい感じでもいいかなっとは思いますね。
	// ただrepository自体が日本語だと保管庫≒永続層って意味なので、repository/user_register/csv.goでもアリかなーとは思います

	// flags
	var idp = flag.String("id-provider", "uuid", "an id provider[uuid random]")
	var be = flag.String("backend", "csv", "a backend that stores user data [csv stdout]")
	flag.Parse()

	// ここまだddd-guysでとりあげてないですが、idProviderFactoryみたいなの作って変数のばらつきを隠蔽してもかもですね
	var idpRegister duser.IIdProvider
	switch *idp {
	case "uuid":
		idpRegister = idprovider.NewUUIDIDProvider()
	case "random":
		idpRegister = idprovider.NewRandomStringIDProvider(32)
	default:
		log.Printf("Invalid id provider [%s]", *idp)
		os.Exit(1)
	}

	// IIdStorer っていう名前があまりイケてない・・
	// backend って使っちゃっているのでわかりにくくなっている可能性がある
	var backendRegister duser.IIdStorer
	switch *be {
	case "csv":
		backendRegister = backend.NewCSVBackend()
	case "stdout":
		backendRegister = backend.NewSTDOUTBackend()
	default:
		log.Printf("Invalid backend [%s]", *be)
		os.Exit(1)
	}

	firstName := "John"
	lastName := "Smith"

	// factoryパターン使うなら
	// usecase.NewUserRegisterUsecase(
	//   idprovider.CreateIdProvider(*idp),
	//   userRepositoryFactory(*be),
	// )
	// とかになるので、スッキリするかもですね
	err := register.NewUserRegisterUsecase(
		idpRegister,
		backendRegister,
	).Execute(
		firstName,
		lastName,
	)
	if err != nil {
		log.Fatalf("Failed to register user. raw error message: %v", err)
	}
	log.Println("Succeeded to register user")
}
