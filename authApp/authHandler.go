package authApp

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/izaakdale/go-auth/dto"
	"github.com/izaakdale/go-auth/service"
)

type AuthHandler struct {
	service service.AuthService
}

func (authHandler AuthHandler) Login(writer http.ResponseWriter, request *http.Request) {

	var loginRequest dto.LoginRequest
	if err := json.NewDecoder(request.Body).Decode(&loginRequest); err != nil {
		log.Println("Error decoding login request")
		writer.WriteHeader(http.StatusBadRequest)
	} else {
		token, err := authHandler.service.Login(loginRequest)
		if err != nil {
			writer.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(writer, err.Error())
		} else {

			json.NewEncoder(writer).Encode(dto.TokenResponse{Token: *token})
		}
	}
}

func (authHandler AuthHandler) Register(writer http.ResponseWriter, request *http.Request) {

	// register functionality
	log.Println("Reaching register")
}

func (authHandler AuthHandler) Verify(writer http.ResponseWriter, request *http.Request) {

	// verify functionality
}
