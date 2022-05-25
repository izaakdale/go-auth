package service

import "github.com/izaakdale/auth/domain"

type AuthService interface {
	// db functions go here
}

type DefaultAuthService struct {
	repo domain.AuthRepo
}

func NewAuthRepoDb(repo domain.AuthRepo) DefaultAuthService {
	return DefaultAuthService{
		repo,
	}
}
