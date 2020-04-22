package user_test

import (
	"testing"

	"github.com/jupemara/ddd-guys/go/domain/model/user"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUserBdd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "create user instance test")
}

var _ = Describe("NewUser", func() {
	var (
		id *user.Id
	)
	BeforeEach(func() {
		id = user.NewId("valid-user-id")
	})
	Context("with valid first name and last name", func() {
		It("should create User instance without error", func() {
			const (
				firstName = "John"
				lastName  = "Smith"
			)
			user, err := user.NewUser(id, firstName, lastName)
			Expect(err).To(BeNil())
			Expect(user.Name().FirstName()).To(Equal(firstName))
			Expect(user.Name().LastName()).To(Equal(lastName))
		})
	})
	Context("with empty first name", func() {
		It("should occur assertion error", func() {
			const (
				firstName = ""
				lastName  = "Smith"
			)
			user, err := user.NewUser(id, firstName, lastName)
			Expect(user).To(BeNil())
			Expect(err.Error()).To(Equal("assertion error"))
		})
	})
	Context("with empty last name", func() {
		It("should occur assertion error", func() {
			const (
				firstName = "John"
				lastName  = ""
			)
			user, err := user.NewUser(id, firstName, lastName)
			Expect(user).To(BeNil())
			Expect(err.Error()).To(Equal("assertion error"))
		})
	})
	Context("with empty both last name and first name", func() {
		It("should occur assertion error", func() {
			const (
				firstName = ""
				lastName  = ""
			)
			user, err := user.NewUser(id, firstName, lastName)
			Expect(user).To(BeNil())
			Expect(err.Error()).To(Equal("assertion error"))
		})
	})
})
