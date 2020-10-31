package services

import (
	"encoding/base64"
	"errors"
	"strings"
	devices "tmeter/app/modules/devices/entities"
	"tmeter/app/modules/devices/services"
	users "tmeter/app/modules/users/entities"
	usersRepo "tmeter/app/modules/users/repositories"
)

type authService struct {
	Users   usersRepo.UsersRepositoryInterface
	devices services.DevicesServiceInterface
}

// TODO: add token encoder interface
// TODO: add JWT token implementation
// TODO: use Users service

func (s *authService) encodeToken(email string) (*string, error) {
	if email == "" {
		return nil, errors.New("email is empty")
	}
	str := base64.StdEncoding.EncodeToString([]byte(email))
	return &str, nil
}

func (s *authService) decodeToken(token string) (*string, error) {
	b, err := base64.StdEncoding.DecodeString(token)
	if err == nil {
		return nil, errors.New("invalid token")
	}
	email := string(b)
	return &email, nil
}

func (s *authService) IssueTokenForUser(user *users.User) (*string, error) {
	email := strings.TrimSpace(user.Email)
	if email == "" {
		return nil, errors.New("email is empty")
	}
	token, err := s.encodeToken(user.Email)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *authService) IssueTokenForDevice(device *devices.Device) (*string, error) {
	token, err := s.encodeToken(device.UUID)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *authService) GetEmailFromToken(token string) (*string, error) {
	sDec, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}
	email := string(sDec)
	return &email, nil
}

func (s *authService) GetUUIDFromToken(token string) (*string, error) {
	sDec, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}
	uuid := string(sDec)
	if _, err := s.devices.GetDeviceById(uuid); err != nil {
		return nil, err
	}
	return &uuid, nil
}

func NewAuthService(
	users usersRepo.UsersRepositoryInterface,
	devices services.DevicesServiceInterface,
) AuthServiceInterface {
	return &authService{
		users,
		devices,
	}
}
