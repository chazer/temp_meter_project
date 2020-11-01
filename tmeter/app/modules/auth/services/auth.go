package services

import (
	"errors"
	"strings"
	devices "tmeter/app/modules/devices/entities"
	"tmeter/app/modules/devices/services"
	users "tmeter/app/modules/users/entities"
	usersRepo "tmeter/app/modules/users/repositories"
	"tmeter/lib/tokens"
)

type authService struct {
	Users   usersRepo.UsersRepositoryInterface
	devices services.DevicesServiceInterface
	encoder tokens.TokenEncoderInterface
}

func (s *authService) IssueTokenForUser(user *users.User) (*string, error) {
	email := strings.TrimSpace(user.Email)
	if email == "" {
		return nil, errors.New("email is empty")
	}
	token, err := s.encoder.EncodeToken(user.Email)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *authService) IssueTokenForDevice(device *devices.Device) (*string, error) {
	token, err := s.encoder.EncodeToken(device.UUID)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *authService) GetEmailFromToken(token string) (*string, error) {
	email, err := s.encoder.DecodeToken(token)
	if err != nil {
		return nil, err
	}
	return email, nil
}

func (s *authService) GetUUIDFromToken(token string) (*string, error) {
	uuid, err := s.encoder.DecodeToken(token)
	if err != nil {
		return nil, err
	}
	if _, err := s.devices.GetDeviceById(*uuid); err != nil {
		return nil, err
	}
	return uuid, nil
}

func NewAuthService(
	users usersRepo.UsersRepositoryInterface,
	devices services.DevicesServiceInterface,
	encoder tokens.TokenEncoderInterface,
) AuthServiceInterface {
	return &authService{
		users,
		devices,
		encoder,
	}
}
