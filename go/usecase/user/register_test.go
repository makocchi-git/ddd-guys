package user_test

import (
	"errors"
	"log"
	"testing"

	"github.com/golang/mock/gomock"
	mock "github.com/jupemara/ddd-guys/go/mock/domain/model/user"
	"github.com/jupemara/ddd-guys/go/usecase/user_register"
)

func TestUserRegister_Execute(t *testing.T) {
	controller := gomock.NewController(t)
	// よくcontroller.Finish()を呼び忘れるw
	defer controller.Finish()
	repository := mock.NewMockIUserRepository(controller)
	usecase := user_register.NewUserRegisterUsecase(repository)
	// mock repositoryが1回呼び出されることを想定
	// mock repositoryがnilを返すことを想定
	// 試しにTimes(1)をTimes(2)にするとテストがこけます
	repository.EXPECT().Store(gomock.Any()).Return(nil).Times(1)
	err := usecase.Execute("John", "Smith")
	if err != nil {
		log.Fatalf("unexpected error occurred: %v", err)
	}
}

func TestUserRegister_Execute_InvalidFirstName(t *testing.T) {
	controller := gomock.NewController(t)
	// よくcontroller.Finish()を呼び忘れるw
	defer controller.Finish()
	repository := mock.NewMockIUserRepository(controller)
	usecase := user_register.NewUserRegisterUsecase(repository)
	// ドメインロジックで先にこけるのでTimes(0)になるはず
	repository.EXPECT().Store(gomock.Any()).Return(nil).Times(0)
	err := usecase.Execute("", "Smith")
	if err == nil {
		log.Fatal("should occurr error. however err is nil")
	}
}

func TestUserRegister_Execute_RepositoryError(t *testing.T) {
	controller := gomock.NewController(t)
	// よくcontroller.Finish()を呼び忘れるw
	defer controller.Finish()
	repository := mock.NewMockIUserRepository(controller)
	usecase := user_register.NewUserRegisterUsecase(repository)
	// repositoryがエラーのケース(例えばDBがなんらかのエラーが起こったと想定)
	repository.EXPECT().Store(gomock.Any()).Return(errors.New("test error")).Times(1)
	err := usecase.Execute("John", "Smith")
	if err == nil {
		log.Fatal("should occurr error. however err is nil")
	}
}
