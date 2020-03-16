package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-openapi/runtime/middleware/header"
	"github.com/jupemara/ddd-guys/go/usecase/user"
	usecase "github.com/jupemara/ddd-guys/go/usecase/user"
)

type registRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// 単純に結果を msg として返す
// 場合によってはオリジナルの return code が増えることもあるかもしれない
type registResponse struct {
	Msg string `json:"msg"`
}

const (
	msgRegistSuccess = "succeeded to regist user"
	msgRegistFail    = "failed to regist user"
)

type HttpUserRegistController struct {
	usecase *usecase.UserRegistUsecase
	output  IOutputPort
}

func NewRegistConroller(usecase *user.UserRegistUsecase, output IOutputPort) *HttpUserRegistController {
	return &HttpUserRegistController{
		usecase: usecase,
		output:  output,
	}
}

func (c *HttpUserRegistController) Register(url string, mux *http.ServeMux) {
	mux.HandleFunc(url, c.HandlerFunc)
}

func (c *HttpUserRegistController) HandlerFunc(
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

	var reqBody registRequest
	if err := body.Decode(&reqBody); err != nil {
		// return error
		http.Error(w, "invalid json format", http.StatusBadRequest)
		return
	}

	firstName := reqBody.FirstName
	lastName := reqBody.LastName
	if err := c.usecase.Execute(firstName, lastName); err != nil {
		http.Error(w, msgRegistFail, http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(registResponse{
		Msg: msgRegistSuccess,
	})
	if err != nil {
		http.Error(w, "unexpected error occurred", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}
