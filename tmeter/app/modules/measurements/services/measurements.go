package services

import "tmeter/app/modules/measurements/repositories"

type MeasurementsServiceInterface interface {
	WriteTemperature(time int64, value float32)
}

type MeasurementsService struct {
	tempLog repositories.FloatLogInterface
}

func NewMeasurementsService(temp repositories.TempLogInterface) *MeasurementsService {
	return &MeasurementsService{
		tempLog: temp,
	}
}

func (s *MeasurementsService) WriteTemperature(time int64, value float32) {
	s.tempLog.Set(time, value)
}
