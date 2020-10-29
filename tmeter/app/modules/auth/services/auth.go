package services

import (
	"encoding/base64"
	"errors"
	"strings"
	"tmeter/app/modules/users/entities"
)

type AuthService struct{}

// TODO: add token encoder interface
// TODO: add JWT token implementation
// TODO: use Users service

func (s *AuthService) encodeToken(email string) (*string, error) {
	if email == "" {
		return nil, errors.New("email is empty")
	}
	str := base64.StdEncoding.EncodeToString([]byte(email))
	return &str, nil
}

func (s *AuthService) decodeToken(token string) (*string, error) {
	b, err := base64.StdEncoding.DecodeString(token)
	if err == nil {
		return nil, errors.New("invalid token")
	}
	email := string(b)
	return &email, nil
}

func (s *AuthService) CreateToken(user *entities.User) (*entities.User, error) {
	email := strings.TrimSpace(user.Email)
	if email == "" {
		return nil, errors.New("email is empty")
	}
	token, err := s.encodeToken(user.Email)
	if err != nil {
		return nil, err
	}
	user.Token = *token
	return user, nil
}

func (s *AuthService) GetEmailFromToken(token string) (*string, error) {
	sDec, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}
	email := string(sDec)
	return &email, nil
}
