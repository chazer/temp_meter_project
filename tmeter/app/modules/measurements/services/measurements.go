package services

import (
	devices "tmeter/app/modules/devices/entities"
	"tmeter/app/modules/devices/services"
	"tmeter/app/modules/measurements/repositories"
)

type MeasurementsServiceInterface interface {
	WriteTemperature(uuid string, time int64, value float32) error
	GetDeviceLog(uuid string) repositories.DeviceLogInterface
}

type MeasurementsService struct {
	logs            map[string]repositories.DeviceLogInterface
	floatLogFactory func() repositories.DeviceLogInterface
	devices         services.DevicesServiceInterface
}

func NewMeasurementsService(
	fn func() repositories.DeviceLogInterface,
	devices services.DevicesServiceInterface,
) *MeasurementsService {
	return &MeasurementsService{
		logs:            make(map[string]repositories.DeviceLogInterface),
		floatLogFactory: fn,
		devices:         devices,
	}
}

func (s *MeasurementsService) WriteTemperature(uuid string, time int64, value float32) error {
	d, err := s.devices.GetDeviceById(uuid)
	if err != nil {
		return err
	}
	s.GetDeviceLog(d.UUID).Append(&devices.DeviceLogRecord{
		Time:        time,
		Temperature: &value,
	})
	s.devices.Touch(uuid)
	return nil
}

func (s *MeasurementsService) GetDeviceLog(uuid string) repositories.DeviceLogInterface {
	m := s.logs[uuid]
	if m == nil {
		m = s.floatLogFactory()
		s.logs[uuid] = m
	}
	return m
}
