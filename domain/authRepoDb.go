package domain

import (
	"database/sql"
	"log"

	"github.com/izaakdale/utils-go/logger"
	"github.com/izaakdale/utils-go/response"
	"github.com/jmoiron/sqlx"
)

type AuthRepoDb struct {
	client sqlx.DB
}

func NewAuthRepoDb(dbClient *sqlx.DB) AuthRepoDb {
	return AuthRepoDb{
		client: *dbClient,
	}
}
func (authRepoDb AuthRepoDb) FindBy(username string, password string) (*Login, *response.ErrorReponse) {

	var login Login

	sqlQuery := `SELECT users.username, users.customer_id, users.role, group_concat(accounts.account_id) as account_numbers from users
				LEFT JOIN accounts on accounts.customer_id = users.customer_id
				WHERE username=? and password=?
				group by users.username, users.customer_id, users.role`

	err := authRepoDb.client.Get(&login, sqlQuery, username, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.NewAuthenticationError("Invalid credentials")
		} else {
			log.Println("Error while verifying login request " + err.Error())
			return nil, response.NewUnexpectedError("Unexpected DB Error")
		}
	}
	return &login, nil
}

func (authRepoDb AuthRepoDb) GenerateAndSaveRefreshToken(authT AuthToken) (string, *response.ErrorReponse) {
	var refreshToken string
	var err *response.ErrorReponse

	if refreshToken, err = authT.newRefreshToken(); err != nil {
		return "", err
	}

	sqlStatement := "INSERT INTO refresh_token_store (refresh_token) values (?)"
	_, sqlErr := authRepoDb.client.Exec(sqlStatement, refreshToken)

	if sqlErr != nil {
		logger.Error("Unexpected DB Error: " + sqlErr.Error())
		return "", response.NewUnexpectedError("Unexpected DB Error")
	}
	return refreshToken, nil

}
