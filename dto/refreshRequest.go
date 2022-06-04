package dto

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/izaakdale/utils-go/logger"
)

type RefreshTokenRequest struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (r RefreshTokenRequest) IsAccessTokenValid() *jwt.ValidationError {

	_, err := jwt.Parse(r.AccessToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("HMAC_SERVER_SECRET")), nil
	})
	if err != nil {
		logger.Error("Error while parsing token " + err.Error())
		var vErr *jwt.ValidationError
		if errors.As(err, vErr) {
			return vErr
		}
	}
	return nil
}
