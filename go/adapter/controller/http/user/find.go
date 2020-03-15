package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jupemara/ddd-guys/go/usecase/user"
)

type response struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type HttpUserFindController struct {
	usecase user.UserFindUsecase
	output  IOutputPort
}

// /users.xml?id=12345
// /users.json?id=12345

type IOutputPort interface {
	Print(dto usecase.Dto) string
}

type JsonOutputPort struct {
}

func (o *JsonOutputPort) Print(dto) string {
	return fmt.Sprintf(`{
id: %s
}`, dto.Id)
}

// 関数名にNewを使った場合
func New(usecase user.UserFindUsecase) *HttpUserFindController {
	return &HttpUserFindController{usecase}
}

func (c *HttpUserFindController) Register(url string) {
	http.HandleFunc(url, c.HandlerFunc)
}

func (c *HttpUserFindController) HandlerFunc(
	w http.ResponseWriter,
	r *http.Request,
) {
	id := r.URL.Query().Get("id")
	if len(id) <= 0 {
		http.Error(w, "no privided id", http.StatusBadRequest)
		return
	}
	dto, err := c.usecase.Execute(id)
	if err != nil {
		http.Error(w, "unexpected error occurred", http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(response{
		Id:        dto.Id,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
	})
	if err != nil {
		http.Error(w, "unexpected error occurred", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}
