package user_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jupemara/ddd-guys/go/adapter/repository/user/mock"

	domain "github.com/jupemara/ddd-guys/go/domain/model/user"
	"github.com/jupemara/ddd-guys/go/usecase/user"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUserUpdateUsecase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "usecase test for updating user name")
}

var _ = Describe("UseUpdateUsecase", func() {
	var (
		ctrl         *gomock.Controller
		repository   *mock.MockIUserRepository
		usecase      *user.UserUpdateUsecase
		existingUser *domain.User
	)
	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		repository = mock.NewMockIUserRepository(ctrl)
		usecase = user.NewUserUpdateUsecase(repository)
		existingUser, _ = domain.NewUser(domain.NewId("12345"), "Rob", "Pike")
	})
	AfterEach(func() {
		ctrl.Finish()
	})
	Describe("#Execute", func() {
		Context("with valid name", func() {
			It("updates an User without error", func() {
				command := user.NewCommand("12345", "Ken", "Thompson")
				repository.EXPECT().FindById(gomock.Any()).Return(existingUser, nil).Times(1)
				repository.EXPECT().Update(gomock.Any()).Return(nil).Times(1)
				err := usecase.Execute(command)
				Expect(err).To(BeNil())
			})
		})
		Context("with empty first name and valid last name", func() {
			It("updates an User without error", func() {
				command := user.NewCommand("12345", "", "Thompson")
				repository.EXPECT().FindById(gomock.Any()).Return(existingUser, nil).Times(1)
				repository.EXPECT().Update(gomock.Any()).Return(nil).Times(1)
				err := usecase.Execute(command)
				Expect(err).To(BeNil())
			})
		})
		Context("with empty first/last name", func() {
			It("does not update User with no error", func() {
				command := user.NewCommand("12345", "", "")
				repository.EXPECT().FindById(gomock.Any()).Return(existingUser, nil).Times(1)
				repository.EXPECT().Update(gomock.Any()).Return(nil).Times(0)
				err := usecase.Execute(command)
				Expect(err).To(BeNil())
			})
		})
		Context("when repository occur unexpected error", func() {
			It("fails with error when executing FindById()", func() {
				command := user.NewCommand("12345", "Ken", "Thompson")
				repository.EXPECT().FindById(gomock.Any()).Return(nil, errors.New("unexpected error")).Times(1)
				repository.EXPECT().Update(gomock.Any()).Return(nil).Times(0)
				err := usecase.Execute(command)
				Expect(err).NotTo(BeNil())
			})
		})
	})
})
