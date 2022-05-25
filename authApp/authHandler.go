package authApp

import (
	"net/http"

	"github.com/izaakdale/auth/service"
)

type AuthHandler struct {
	service service.AuthService
}

func (authHandler AuthHandler) Login(writer http.ResponseWriter, request *http.Request) {

	// login functionality
}

func (authHandler AuthHandler) Register(writer http.ResponseWriter, request *http.Request) {

	// register functionality
}

func (authHandler AuthHandler) Verify(writer http.ResponseWriter, request *http.Request) {

	// verify functionality
}
