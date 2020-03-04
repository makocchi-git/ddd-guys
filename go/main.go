package main

import (
	"flag"
	"log"
	"os"

	repository "github.com/jupemara/ddd-guys/go/adapter/repository/user"
	usecase "github.com/jupemara/ddd-guys/go/usecase/user"
)

func register(firstName, lastName string) {
	// ここでは直接アプリケーションサービス(ユースケース)を呼び出していますが、実際はCLIがあったりHTTPのハンドラーがあって、それらからfirstName, lastNameを取得する形になります
	// 技術的複雑性は"すべて"DI(Dependency Injection: 依存性の注入)を使って外側から入れます
	err := usecase.NewUserRegisterUsecase(
		repository.NewUuidIdProvider(),
		repository.NewCsvRepository(),
	).Execute(
		firstName,
		lastName,
	)
	if err != nil {
		log.Fatalf("Failed to register user. raw error message: %v", err)
	}
	log.Printf("Succeeded to register user [%s %s]\n", firstName, lastName)
}

func main() {
	const (
		newFirstName = "Taro"
		newLastName  = "Yamada"
	)

	// flag
	var (
		gen = flag.Bool("gen-user", false, "generate sample user data")
		id  = flag.String("id", "", "target id for modify user name")
	)
	flag.Parse()

	// generate user data and exit
	if *gen {
		register("John", "Doh")
		register("Rob", "Pike")
		register("Ken", "Thompson")
		os.Exit(0)
	}

	if len(*id) < 1 {
		log.Fatal("please set id for modify user name")
	}

	command := usecase.NewCommand(
		*id,
		newFirstName,
		newLastName,
	)

	err := usecase.NewUserUpdateUsecase(
		repository.NewCsvRepository(),
	).Execute(command)
	if err != nil {
		log.Fatalf("unexepcted error occurred: %s", err)
	}

	log.Printf("Succeeded to change name to %s %s\n", newFirstName, newLastName)
}
