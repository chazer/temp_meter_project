package services

import (
	"errors"
	"time"
	dr "tmeter/app/modules/devices/repositories"
)
import "tmeter/app/modules/devices/entities"

const ErrNoDevice = "no devices"

type DevicesServiceInterface interface {
	CreateDevice(name string, email string) (*entities.Device, error)
	GetDeviceById(uuid string) (*entities.Device, error)
	// TODO: return Cursor
	GetDevicesByEmail(email string) ([]*entities.Device, error)
	Touch(uuid string)
	GetAllDevices() []*entities.Device
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

func (s *DevicesService) CreateDevice(name string, email string) (*entities.Device, error) {
	d := entities.MakeDevice()
	d.Name = name
	d.OwnerEmail = email
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

func (s *DevicesService) Touch(uuid string) {
	device, _ := s.GetDeviceById(uuid)
	if device != nil {
		device.UpdatedAt = time.Now()
		s.registry.Replace(uuid, device)
	}
}

func (s *DevicesService) GetAllDevices() []*entities.Device {
	return s.registry.Items()
}
