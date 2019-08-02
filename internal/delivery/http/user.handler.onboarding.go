package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zainul/ark/xsend"
	"github.com/zainul/arkana-kit/internal/contract"
	"github.com/zainul/arkana-kit/internal/pkg/error/deliveryerror"
	"github.com/zainul/arkana-kit/internal/pkg/error/usecaseerror"
	"github.com/zainul/arkana-kit/internal/usecase"
	"github.com/zainul/nux"
)

// UserHandler ...
type UserHandler struct {
	UserUsecase usecase.User
}

// ActivateUser ...
func (u *UserHandler) ActivateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	response := contract.Response{}
	bodyReq := contract.ActivateUser{}

	vars := mux.Vars(r)
	code := vars["code"]

	if code == "" {
		response.Error = deliveryerror.GetError(usecaseerror.MissingRequiredParam, errors.New("missing code"))
		xsend.Write(w, response, http.StatusInternalServerError)
		return
	}

	bodyReq.Code = code

	errUsecase := u.UserUsecase.ActivateUser(ctx, bodyReq)

	if errUsecase != nil {
		response.Error = errUsecase
		xsend.Write(w, response, http.StatusInternalServerError)
		return
	}

	xsend.Write(w, response)
	return

}

// RegisterUser ...
func (u *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	response := contract.Response{}
	bodyReq := contract.RegisterUser{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&bodyReq)

	if err != nil {
		response.Error = deliveryerror.GetError(usecaseerror.MissingRequiredParam, err)
		xsend.Write(w, response, http.StatusInternalServerError)
		return
	}

	listError := nux.ValidateStruct(bodyReq)

	if len(listError) > 0 {
		response.Error = deliveryerror.GetError(
			listError[0].Code,
			errors.New(listError[0].Message),
		)
		xsend.Write(w, response, http.StatusInternalServerError)
		return
	}

	errRegister := u.UserUsecase.RegisterUser(ctx, bodyReq)

	if errRegister != nil {
		response.Error = errRegister
		xsend.Write(w, response, http.StatusInternalServerError)
		return
	}

	xsend.Write(w, response)
	return
}
