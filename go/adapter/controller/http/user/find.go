package user

import (
	"encoding/json"
	"net/http"

	"github.com/jupemara/ddd-guys/go/usecase/user"
	usecase "github.com/jupemara/ddd-guys/go/usecase/user"
)

type findResponse struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type HttpUserFindController struct {
	usecase *usecase.UserFindUsecase
	// TODO:
	// output  IOutputPort
}

// 関数名にNewを使った場合
func NewFindController(usecase *user.UserFindUsecase) *HttpUserFindController {
	return &HttpUserFindController{
		usecase: usecase,
	}
}

func (c *HttpUserFindController) Register(url string, mux *http.ServeMux) {
	mux.HandleFunc(url, c.HandlerFunc)
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
	res, err := json.Marshal(findResponse{
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
