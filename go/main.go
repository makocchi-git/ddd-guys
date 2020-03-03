package main

import (
	"log"

	repository "github.com/jupemara/ddd-guys/go/adapter/repository/user"
	usecase "github.com/jupemara/ddd-guys/go/usecase/user"
)

func register() {
	firstName := "John"
	lastName := "Smith"
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
	log.Println("Succeeded to register user")
}

func main() {
	const id = "abe7da64-5d25-11ea-aa98-acde48001122"
	user, err := usecase.NewUserFindUsecase(
		repository.NewCsvRepository(),
	).Execute(id)
	if err != nil {
		log.Fatalf("unexepcted error occurred: %s", err)
	}
	log.Printf(`id: %s
first name: %s
last name: %s`, user.Id, user.FirstName, user.LastName)
}
