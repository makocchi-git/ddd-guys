package main

import (
	"flag"
	"log"
	"os"

	duser "github.com/makocchi-git/ddd-guys/go/pkg/domain/user"

	backend "github.com/makocchi-git/ddd-guys/go/pkg/repository/user_register/backend"
	idprovider "github.com/makocchi-git/ddd-guys/go/pkg/repository/user_register/id_provider"

	register "github.com/makocchi-git/ddd-guys/go/pkg/usecase/user_register"
)

func main() {
	// go っぽく pkg/ にいろいろ入れてみた
	// 個人的にはこういう構成の方が見やすいかも
	// でも package 名は苦労しますな・・
	// 今回は domain/user にして domainuser という意味で "duser" にしてある
	// usecase の package 名は usecase の u を接頭語としてみたけども・・

	// 試しに Random String で生成する IDP と保存するんじゃなくて STDOUT に出力する Store() を作ってみた
	// repository の directory 構成的に .../backend と .../id_provier みたいに細かくしたんだけど
	// そういうのはアリなのかな？

	// flags
	var idp = flag.String("id-provider", "uuid", "an id provider[uuid random]")
	var be = flag.String("backend", "csv", "a backend that stores user data [csv stdout]")
	flag.Parse()

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
