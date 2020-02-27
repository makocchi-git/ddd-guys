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
	idprovider "github.com/makocchi-git/ddd-guys/go/pkg/repository/user/register/id/provider"
	// パッケージエイリアスは愚直にusecaseのがわかりやすいかもですね
	// >> ここも同じく usecase/user/register にしてみる
	//    domain と同じように複数の usecase が出てきた時に分ければいいので、ここでは usecase で定義
	usecase "github.com/makocchi-git/ddd-guys/go/pkg/usecase/user/register"
	// >> まとめると
	//    pkg/domain/user/...
	//    pkg/repository/user/...
	//    pkg/usecase/user/...
	//    という風になり、わかりやすくなったかな？
	//    むしろこういう風に合わせちゃうと、なんかしらの縛りっぽくなっちゃうのならば逆効果だけども・・
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
