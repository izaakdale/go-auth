package domain

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/izaakdale/utils-go/logger"
	"github.com/izaakdale/utils-go/response"
)

type AuthToken struct {
	token *jwt.Token
}

func (t AuthToken) NewAccessToken() (string, *response.ErrorReponse) {
	signedTokenStr, err := t.token.SignedString([]byte(os.Getenv("HMAC_SERVER_SECRET")))
	if err != nil {
		logger.Error("Error signing token with env secret " + err.Error())
		return "", response.NewUnexpectedError("Token generate failed")
	}
	return signedTokenStr, nil
}

func NewAuthToken(claims AccessTokenClaims) AuthToken {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return AuthToken{
		token: token,
	}
}
