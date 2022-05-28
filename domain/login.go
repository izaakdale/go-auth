package domain

import (
	"database/sql"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/izaakdale/utils-go/logger"
	"github.com/izaakdale/utils-go/response"
)

const TOKEN_DURATION = time.Hour

type Login struct {
	Username   string         `db:"username"`
	CustomerId sql.NullString `db:"customer_id"`
	Accounts   sql.NullString `db:"account_numbers"`
	Role       string         `db:"role"`
}

func (login Login) GenerateToken() (*string, *response.ErrorReponse) {

	// define a map of type string to interface, meaning you can define keys to store whatever
	var claims jwt.MapClaims
	// always check sql.NullString for validity
	if login.CustomerId.Valid && login.Accounts.Valid {
		claims = login.claimsForUser()
	} else {
		claims = login.claimsForAdmin()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedTokenStr, err := token.SignedString([]byte(os.Getenv("HMAC_SERVER_SECRET")))
	if err != nil {
		logger.Error("Error signing token with env secret " + err.Error())
		return nil, response.NewUnexpectedError("Token generate failed")
	}
	return &signedTokenStr, nil
}

func (login Login) claimsForUser() jwt.MapClaims {
	accounts := strings.Split(login.Accounts.String, ",")
	return jwt.MapClaims{
		"username":    login.Username,
		"customer_id": login.CustomerId.String,
		"accounts":    accounts,
		"role":        login.Role,
		"exp":         time.Now().Add(TOKEN_DURATION).Unix(),
	}
}
func (login Login) claimsForAdmin() jwt.MapClaims {
	return jwt.MapClaims{
		"role":     login.Role,
		"username": login.Username,
		"exp":      time.Now().Add(TOKEN_DURATION).Unix(),
	}
}
