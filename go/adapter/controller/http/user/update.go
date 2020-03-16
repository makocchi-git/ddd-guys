package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-openapi/runtime/middleware/header"
	"github.com/jupemara/ddd-guys/go/usecase/user"
	usecase "github.com/jupemara/ddd-guys/go/usecase/user"
)

type updateRequest struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// 単純に結果を msg として返す
// 場合によってはオリジナルの return code が増えることもあるかもしれない
type updateResponse struct {
	Msg string `json:"msg"`
}

const (
	msgUpdateSuccess = "succeeded to update user"
	msgUpdateFail    = "failed to update user"
)

type HttpUserUpdateController struct {
	usecase *usecase.UserUpdateUsecase
	output  IOutputPort
}

func NewUpdateController(usecase *user.UserUpdateUsecase, output IOutputPort) *HttpUserUpdateController {
	return &HttpUserUpdateController{
		usecase: usecase,
		output:  output,
	}
}

func (c *HttpUserUpdateController) Register(url string, mux *http.ServeMux) {
	mux.HandleFunc(url, c.HandlerFunc)
}

func (c *HttpUserUpdateController) HandlerFunc(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Header check
	if r.Header.Get("Content-Type") != "" {
		h, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if h != "application/json" {
			http.Error(w, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
			return
		}
	}

	body := json.NewDecoder(r.Body)
	body.DisallowUnknownFields()

	var reqBody updateRequest
	if err := body.Decode(&reqBody); err != nil {
		// return error
		http.Error(w, "invalid json format", http.StatusBadRequest)
		return
	}

	id := reqBody.Id
	firstName := reqBody.FirstName
	lastName := reqBody.LastName
	command := usecase.NewCommand(id, firstName, lastName)
	if err := c.usecase.Execute(command); err != nil {
		http.Error(w, msgRegistFail, http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(registResponse{
		Msg: msgUpdateSuccess,
	})
	if err != nil {
		http.Error(w, "unexpected error occurred", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}
