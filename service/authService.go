package service

import (
	"log"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/izaakdale/go-auth/domain"
	"github.com/izaakdale/go-auth/dto"
)

type AuthService interface {
	Login(request dto.LoginRequest) (*string, error)
	Verify(urlParams map[string]string) (bool, error)
}
type DefaultAuthService struct {
	repo domain.AuthRepoDb
}

func NewAuthRepoDb(repo domain.AuthRepoDb) DefaultAuthService {
	return DefaultAuthService{
		repo,
	}
}

func (authService DefaultAuthService) Login(request dto.LoginRequest) (*string, error) {

	login, err := authService.repo.FindBy(request.Username, request.Password)
	if err != nil {
		return nil, err
	}
	claims, err := login.GenerateToken()

	return claims, nil
}

func (authService DefaultAuthService) Verify(urlParams map[string]string) (bool, error) {

	if jwtToken, err := jwtTokenFromString(urlParams["Authorization"]); err != nil {
		return false, err
	} else {
		if jwtToken.Valid {
			claims := jwtToken.Claims.(*domain.AccessTokenClaims)

			log.Println(claims.Role)

		} else {
			log.Println("Invalid token")
		}
	}

	return true, nil
}

func jwtTokenFromString(fullString string) (*jwt.Token, error) {

	splitString := strings.Split(fullString, "Bearer ")
	tokenString := splitString[1]

	token, err := jwt.ParseWithClaims(tokenString, &domain.AccessTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("HMAC_SERVER_SECRET")), nil
	})
	if err != nil {
		log.Println("Error parsing token: " + err.Error())
		return nil, err
	}
	return token, nil
}
