package user

import (
	"log"
	"os"

	"github.com/jupemara/ddd-guys/go/usecase/user"
)

type EnvUserFindController struct {
	usecase user.UserFindUsecase
}

// ここはsimpleに New だけでもいいかなっとは思います
// golangのcookiejarパッケージでは実際にNewを使っています
func NewEnvUserFindController(usecase user.UserFindUsecase) *EnvUserFindController {
	return &EnvUserFindController{usecase}
}

func (c *EnvUserFindController) Handle() {
	id := os.Getenv("DDD_GUYS_USER_ID")
	if len(id) <= 0 {
		log.Fatalf(`no "DDD_GUYS_USER_ID" environment variable...`)
	}
	dto, err := c.usecase.Execute(id)
	if err != nil {
		log.Fatalf(`unexpected error occurred. error: %s`, err)
	}
	log.Printf(`---found user---
id: %s
first_name: %s
last_name: %s
`, dto.Id, dto.FirstName, dto.LastName)
}
