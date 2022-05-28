package domain

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	CustomerId string   `json:"customer_id"`
	Accounts   []string `json:"accounts"`
	Username   string   `json:"username"`
	Expiry     int64    `json:"expiry"`
	Role       string   `json:"role"`
}

type AccessTokenClaims struct {
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
