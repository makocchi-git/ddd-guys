package idprovider

import (
	"math/rand"
	"time"

	duser "github.com/makocchi-git/ddd-guys/go/pkg/domain/user"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type RandomStringIDProvider struct {
	n int
}

func NewRandomStringIDProvider(n int) *RandomStringIDProvider {
	return &RandomStringIDProvider{
		n: n,
	}
}

func (u *RandomStringIDProvider) NextIdentity() (*duser.Id, error) {
	b := make([]byte, u.n)
	for i := range b {
		b[i] = Letters[rand.Intn(len(Letters))]
	}
	return duser.NewId(string(b)), nil
}
