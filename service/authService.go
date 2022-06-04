package service

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/izaakdale/go-auth/domain"
	"github.com/izaakdale/go-auth/dto"
	"github.com/izaakdale/utils-go/logger"
	"github.com/izaakdale/utils-go/response"
)

type AuthService interface {
	Login(request dto.LoginRequest) (*dto.LoginResponse, *response.ErrorReponse)
	Verify(urlParams map[string]string) *response.ErrorReponse
}
type DefaultAuthService struct {
	repo            domain.AuthRepoDb
	rolePermissions domain.RolePermissions
}

func NewAuthRepoDb(repo domain.AuthRepoDb) DefaultAuthService {
	return DefaultAuthService{
		repo,
		domain.GetRolePermissions(),
	}
}

func (authService DefaultAuthService) Login(request dto.LoginRequest) (*dto.LoginResponse, *response.ErrorReponse) {

	var login *domain.Login
	var err *response.ErrorReponse

	if login, err = authService.repo.FindBy(request.Username, request.Password); err != nil {
		logger.Error("Error logging in with username and password: " + err.Message)
		return nil, err
	}

	claims := login.ClaimsForAccessToken()
	authToken := domain.NewAuthToken(claims)

	var accessToken string
	if accessToken, err = authToken.NewAccessToken(); err != nil {
		return nil, err
	}

	refreshToken, err := authService.repo.GenerateAndSaveRefreshToken(authToken)
	if err != nil {
		return nil, err
	}
	return &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (authService DefaultAuthService) Verify(urlParams map[string]string) *response.ErrorReponse {

	if jwtToken, err := jwtTokenFromString(urlParams["token"]); err != nil {
		return err
	} else {
		if jwtToken.Valid {
			claims := jwtToken.Claims.(*domain.AccessTokenClaims)

			if claims.IsUserRole() {
				if !claims.IsValidUserRequest(urlParams) {
					return response.NewForbiddenError("Invalid request for user claims")
				}
			}

			// check permisions
			if !authService.rolePermissions.IsAuthorizedFor(claims.Role, urlParams["routeName"]) {
				return response.NewForbiddenError(fmt.Sprintf("%s is unauthorized", claims.Role))
			}

		} else {
			logger.Error("Invalid token")
			return response.NewForbiddenError("Invalid token")
		}
	}

	return nil
}

func jwtTokenFromString(tokenString string) (*jwt.Token, *response.ErrorReponse) {

	token, err := jwt.ParseWithClaims(tokenString, &domain.AccessTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("HMAC_SERVER_SECRET")), nil
	})
	if err != nil {
		logger.Error("Error parsing token: " + err.Error())
		return nil, response.NewUnexpectedError("Error parsing token: " + err.Error())
	}
	return token, nil
}
