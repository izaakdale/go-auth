package domain

import "github.com/izaakdale/utils-go/response"

type AuthRepo interface {
	FindBy(string, string) (*Login, error)
	GenerateAndSaveRefreshToken(AuthToken) (string, *response.ErrorReponse)
}
