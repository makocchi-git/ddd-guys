package user

import (
	"encoding/json"
	"net/http"

	util "github.com/jupemara/ddd-guys/go/adapter/controller/http/util"
	"github.com/jupemara/ddd-guys/go/usecase/user"
	usecase "github.com/jupemara/ddd-guys/go/usecase/user"
)

type registerRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// 単純に結果を msg として返す
// 場合によってはオリジナルの return code が増えることもあるかもしれない
type registerResponse struct {
	Msg string `json:"msg"`
}

const (
	msgRegisterSuccess = "succeeded in registering user"
	msgRegisterFail    = "failed to register user"
)

type HttpUserRegisterController struct {
	usecase *usecase.UserRegisterUsecase
	// TODO:
	// output  IOutputPort
}

func NewRegisterController(usecase *user.UserRegisterUsecase) *HttpUserRegisterController {
	return &HttpUserRegisterController{
		usecase: usecase,
	}
}

func (c *HttpUserRegisterController) Register(url string, mux *http.ServeMux) {
	mux.HandleFunc(url, c.HandlerFunc)
}

func (c *HttpUserRegisterController) HandlerFunc(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Header check
	if err := util.ValidateContentTypeApplicationJSON(r.Header); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	body := json.NewDecoder(r.Body)
	body.DisallowUnknownFields()

	var reqBody registerRequest
	if err := body.Decode(&reqBody); err != nil {
		// return error
		http.Error(w, "invalid json format", http.StatusBadRequest)
		return
	}

	firstName := reqBody.FirstName
	lastName := reqBody.LastName
	if err := c.usecase.Execute(firstName, lastName); err != nil {
		http.Error(w, msgRegisterFail, http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(registerResponse{
		Msg: msgRegisterSuccess,
	})
	if err != nil {
		http.Error(w, "unexpected error occurred", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}
