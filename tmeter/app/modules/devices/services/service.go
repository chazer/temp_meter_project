package services

import (
	"errors"
	dr "tmeter/app/modules/devices/repositories"
)
import "tmeter/app/modules/devices/entities"

const ErrNoDevice = "no devices"

type DevicesServiceInterface interface {
	CreateDevice(email string) (*entities.Device, error)
	GetDeviceById(uuid string) (*entities.Device, error)
	// TODO: return Cursor
	GetDevicesByEmail(email string) ([]*entities.Device, error)
}

type DevicesService struct {
	registry dr.DevicesRepositoryInterface
}

type DevicesServiceConfig struct {
	Repository dr.DevicesRepositoryInterface
}

func MakeDevicesService(config *DevicesServiceConfig) *DevicesService {
	s := &DevicesService{}
	if s.registry = config.Repository; s.registry == nil {
		s.registry = dr.MakeDevicesInmemoryRepository()
	}
	return s
}

func (s *DevicesService) CreateDevice(email string) (*entities.Device, error) {
	d := entities.MakeDevice()
	d.UserEmail = email
	inserted := s.registry.Insert(&d)
	return inserted, nil
}

func (s *DevicesService) GetDeviceById(uuid string) (*entities.Device, error) {
	d := s.registry.FindByUUID(uuid)
	if d == nil {
		return nil, errors.New(ErrNoDevice)
	}
	return d, nil
}

func (s *DevicesService) GetDevicesByEmail(email string) ([]*entities.Device, error) {
	return s.registry.FindByEmail(email), nil
}
