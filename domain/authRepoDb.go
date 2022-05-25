package domain

import "github.com/jmoiron/sqlx"

type AuthRepoDb struct {
	client sqlx.DB
}

func NewAuthRepoDb(dbClient *sqlx.DB) AuthRepoDb {
	return AuthRepoDb{
		client: *dbClient,
	}
}
