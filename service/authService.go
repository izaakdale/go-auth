package service

import (
	"github.com/izaakdale/go-auth/domain"
	"github.com/izaakdale/go-auth/dto"
)

type AuthService interface {
	Login(request dto.LoginRequest) (*string, error)
}

type DefaultAuthService struct {
	repo domain.AuthRepoDb
}

func (authService DefaultAuthService) Login(request dto.LoginRequest) (*string, error) {

	login, err := authService.repo.FindBy(request.Username, request.Password)
	if err != nil {
		return nil, err
	}
	claims, err := login.GenerateToken()

	return claims, nil
}

func NewAuthRepoDb(repo domain.AuthRepoDb) DefaultAuthService {
	return DefaultAuthService{
		repo,
	}
}
