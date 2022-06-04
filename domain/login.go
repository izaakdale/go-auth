package domain

import (
	"database/sql"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const TOKEN_DURATION = time.Hour

type Login struct {
	Username   string         `db:"username"`
	CustomerId sql.NullString `db:"customer_id"`
	Accounts   sql.NullString `db:"account_numbers"`
	Role       string         `db:"role"`
}

func (login Login) ClaimsForAccessToken() AccessTokenClaims {
	if login.CustomerId.Valid && login.Accounts.Valid {
		return login.claimsForUser()
	} else {
		return login.claimsForAdmin()
	}
}

func (login Login) claimsForUser() AccessTokenClaims {
	accounts := strings.Split(login.Accounts.String, ",")
	return AccessTokenClaims{
		Username:   login.Username,
		CustomerId: login.CustomerId.String,
		Accounts:   accounts,
		Role:       login.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TOKEN_DURATION).Unix(),
		},
	}
}
func (login Login) claimsForAdmin() AccessTokenClaims {
	return AccessTokenClaims{
		Role:     login.Role,
		Username: login.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TOKEN_DURATION).Unix(),
		},
	}
}
