package services

import (
	devices "tmeter/app/modules/devices/entities"
	users "tmeter/app/modules/users/entities"
)

type AuthServiceInterface interface {
	IssueTokenForUser(user *users.User) (*string, error)
	IssueTokenForDevice(device *devices.Device) (*string, error)
	GetEmailFromToken(token string) (*string, error)
	GetUUIDFromToken(token string) (*string, error)
}
