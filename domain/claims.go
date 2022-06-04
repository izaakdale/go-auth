package domain

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const ACCESS_TOKEN_DURATION = time.Hour
const REFRESH_TOKEN_DURATION = time.Hour * 24 * 28

type AccessTokenClaims struct {
	CustomerId string   `json:"customer_id"`
	Accounts   []string `json:"accounts"`
	Username   string   `json:"username"`
	Role       string   `json:"role"`
	jwt.StandardClaims
}

type RefreshTokenClaims struct {
	TokenType  string   `json:"token_type"`
	CustomerId string   `json:"customer_id"`
	Accounts   []string `json:"accounts"`
	Username   string   `json:"username"`
	Role       string   `json:"role"`
	jwt.StandardClaims
}

func (atc AccessTokenClaims) IsUserRole() bool {
	return atc.Role == "user"
}

func (atc AccessTokenClaims) IsValidUserRequest(urlParams map[string]string) bool {
	return atc.IsUserCustomerIdValid(urlParams["customer_Id"]) && atc.IsAccountIdValid(urlParams["account_id"])
}

func (atc AccessTokenClaims) IsUserCustomerIdValid(urlParamCustomerId string) bool {
	return atc.CustomerId != urlParamCustomerId
}

func (atc AccessTokenClaims) IsAccountIdValid(urlParamAccountId string) bool {
	// check that claims account id list matches the url request account id
	if urlParamAccountId != "" {
		// don't need key just string value
		for _, a := range atc.Accounts {
			if a == urlParamAccountId {
				return true
			}
		}
	}
	return false
}

func (c AccessTokenClaims) RefreshTokenClaims() RefreshTokenClaims {
	return RefreshTokenClaims{
		TokenType:  "refresh_token",
		CustomerId: c.CustomerId,
		Accounts:   c.Accounts,
		Username:   c.Username,
		Role:       c.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(REFRESH_TOKEN_DURATION).Unix(),
		},
	}
}

func (c RefreshTokenClaims) AccessTokenClaims() AccessTokenClaims {
	return AccessTokenClaims{
		CustomerId: c.CustomerId,
		Accounts:   c.Accounts,
		Username:   c.Username,
		Role:       c.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ACCESS_TOKEN_DURATION).Unix(),
		},
	}
}
