package service

import (
	"session-22/jwt"
	"session-22/model"
)

type AuthService interface {
	Login(auth model.AuthRequest) (model.AuthResponse, *model.ErrorResponse)
}

type authService struct {
}

func NewAuthService() AuthService {
	return &authService{}
}

func (a authService) Login(auth model.AuthRequest) (model.AuthResponse, *model.ErrorResponse) {

	if auth.Username != auth.Password {
		return model.AuthResponse{}, &model.InvalidCredential
	}

	token, err := jwt.Create(auth.Username)

	if err != nil {
		return model.AuthResponse{}, &model.UnAuthorized
	}

	authResp := model.AuthResponse{Token: token}

	return authResp, nil

}
