package main

import (
	"log"

	repository "github.com/jupemara/ddd-guys/go/adapter/repository/user"
	usecase "github.com/jupemara/ddd-guys/go/usecase/user_register"
)

func main() {
	firstName := "John"
	lastName := "Smith"
	// ここでは直接アプリケーションサービス(ユースケース)を呼び出していますが、実際はCLIがあったりHTTPのハンドラーがあって、それらからfirstName, lastNameを取得する形になります
	// 技術的複雑性は"すべて"DI(Dependency Injection: 依存性の注入)を使って外側から入れます
	err := usecase.NewUserRegisterUsecase(repository.NewCsvRepository()).Execute(
		firstName,
		lastName,
	)
	if err != nil {
		log.Fatalf("Failed to register user. raw error message: %v", err)
	}
	log.Println("Succeeded to register user")
}
