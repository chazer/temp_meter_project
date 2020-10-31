package services

import (
	devices "tmeter/app/modules/devices/entities"
	"tmeter/app/modules/measurements/repositories"
)

type MeasurementsServiceInterface interface {
	WriteTemperature(d *devices.Device, time int64, value float32)
	GetDeviceLog(d *devices.Device) repositories.DeviceLogInterface
}

type MeasurementsService struct {
	logs            map[string]repositories.DeviceLogInterface
	floatLogFactory func() repositories.DeviceLogInterface
}

func NewMeasurementsService(fn func() repositories.DeviceLogInterface) *MeasurementsService {
	return &MeasurementsService{
		logs:            make(map[string]repositories.DeviceLogInterface),
		floatLogFactory: fn,
	}
}

func (s *MeasurementsService) WriteTemperature(d *devices.Device, time int64, value float32) {
	s.GetDeviceLog(d).Append(&devices.DeviceLogRecord{
		Time:        time,
		Temperature: &value,
	})
}

func (s *MeasurementsService) GetDeviceLog(d *devices.Device) repositories.DeviceLogInterface {
	m := s.logs[d.UUID]
	if m == nil {
		m = s.floatLogFactory()
		s.logs[d.UUID] = m
	}
	return m
}
