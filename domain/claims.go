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

func BuildClaimsFromJwt(mapClaims jwt.MapClaims) (*Claims, error) {
	// bytes, err := json.Marshal(mapClaims)
	// if err != nil {
	// 	fmt.Println("Error marshalling jwt map claims")
	// 	return nil, err
	// }

	// return &Claims{
	// 	CustomerId: ,
	// }, nil
	return nil, nil
}
