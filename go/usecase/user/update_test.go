package user_test

import (
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
		ctrl       *gomock.Controller
		repository domain.IUserRepository
		usecase    *user.UserUpdateUsecase
	)
	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		repository = mock.NewMockIUserRepository(ctrl)
		usecase = user.NewUserUpdateUsecase(repository)
	})
	AfterEach(func() {
		ctrl.Finish()
	})
	Describe("#Execute", func() {
		Context("with valid name", func() {})
		Context("with empty first name", func() {})
		Context("when repository occur unexpected error", func() {})
	})
})
