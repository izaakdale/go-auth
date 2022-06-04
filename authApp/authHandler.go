package authApp

import (
	"encoding/json"
	"net/http"

	"github.com/izaakdale/go-auth/dto"
	"github.com/izaakdale/go-auth/service"
	"github.com/izaakdale/utils-go/logger"
	"github.com/izaakdale/utils-go/response"
)

type AuthHandler struct {
	service service.AuthService
}

func (authHandler AuthHandler) Login(writer http.ResponseWriter, request *http.Request) {

	var loginRequest dto.LoginRequest
	if err := json.NewDecoder(request.Body).Decode(&loginRequest); err != nil {
		logger.Error("Error decoding login request")
		response.WriteJson(writer, http.StatusBadRequest, nil)
	} else {
		token, err := authHandler.service.Login(loginRequest)
		if err != nil {
			response.WriteJson(writer, err.Code, err.AsMessage())
		} else {
			response.WriteJson(writer, http.StatusOK, dto.LoginResponse{AccessToken: token.AccessToken})
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
		err := authHandler.service.Verify(urlParams)
		if err != nil {
			logger.Error("Error verifying in service")
			response.WriteJson(writer, err.Code, dto.NotVerifiedReponse{
				IsAuthorized: false,
				Message:      err.Message,
			})
		} else {
			response.WriteJson(writer, http.StatusOK, dto.VerifiedReponse{
				IsAuthorized: true,
			})
		}
	} else {
		response.WriteJson(writer, http.StatusForbidden, dto.NotVerifiedReponse{
			IsAuthorized: false,
			Message:      "Missing token",
		})
	}

}
