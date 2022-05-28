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
}

func (authHandler AuthHandler) Verify(writer http.ResponseWriter, request *http.Request) {

	urlParams := make(map[string]string)

	for k := range request.URL.Query() {
		urlParams[k] = request.URL.Query().Get(k)
	}

	if urlParams["token"] != "" {
		verified, err := authHandler.service.Verify(urlParams)
		if err != nil {
			log.Println("Error verifying in service " + err.Error())
		}
		json.NewEncoder(writer).Encode(dto.VerifyReponse{
			IsAuthorized: verified,
		})
	}

}
