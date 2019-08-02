package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zainul/arkana-kit/internal/usecase"
)

// NewUserHandler ...
func NewUserHandler(route *mux.Router, user usecase.User) {

	handler := UserHandler{
		UserUsecase: user,
	}

	route.HandleFunc("/user/register", handler.RegisterUser).Methods(http.MethodPost)
	route.HandleFunc("/user/activation/{code}", handler.ActivateUser).Methods(http.MethodGet)
}
