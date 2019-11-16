package services

import (
	"errors"

	"github.com/ehardi19/bitment-backend/domain"
)

func (s *Services) Authenticate(email, password string) (*domain.User, error) {
	user, err := s.Repository.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("authentication failed")
	}

	if ok := s.Authentication.CheckPasswordHash(user.Password, password); ok {
		return &user, nil
	}

	return nil, errors.New("authentication failed")
}

func (s *Services) Login(req domain.LoginRequest) (domain.LoginResponse, error) {
	user, err := s.Authenticate(req.Email, req.Password)
	if err != nil {
		return domain.LoginResponse{}, err
	}

	token, err := s.Authentication.GenerateToken(*user)
	if err != nil {
		return domain.LoginResponse{}, err
	}

	return domain.LoginResponse{
		Token: token,
		ID:    user.ID,
	}, nil
}
