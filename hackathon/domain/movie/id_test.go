package movie_test

import (
	"testing"

	"github.com/jupemara/ddd-guys/hackathon/domain/movie"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIdBdd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "create id instance test")
}

var _ = Describe("NewId", func() {
	Context("with empty id", func() {
		It("fails to create Id and returns error", func() {
			idValue := ""
			id, err := movie.NewId(idValue)
			Expect(id).To(BeNil())
			Expect(err).NotTo(BeNil())
		})
	})
	Context("with valid id", func() {
		It("creates Id without error", func() {
			idValue := "valid-id"
			id, err := movie.NewId(idValue)
			Expect(id).NotTo(BeNil())
			Expect(err).To(BeNil())
			Expect(id.Value()).To(Equal(idValue))
		})
	})
	Context("with an id shorter than the criteria", func() {
		It("fails to create Id and returns error", func() {
			idValue := "short"
			id, err := movie.NewId(idValue)
			Expect(id).To(BeNil())
			Expect(err).NotTo(BeNil())
		})
	})
	Context("with an id longer than the criteria", func() {
		It("fails to create Id and returns error", func() {
			idValue := "longId-abcdefghijklmnopqrstuvwxyz"
			id, err := movie.NewId(idValue)
			Expect(id).To(BeNil())
			Expect(err).NotTo(BeNil())
		})
	})
})
