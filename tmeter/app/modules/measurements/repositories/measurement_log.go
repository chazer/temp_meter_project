package repositories

import (
	"sort"
	"tmeter/app/modules/devices/entities"
)

type MeasurementsLog struct {
	records map[int64]*entities.DeviceLogRecord
}

func (s *MeasurementsLog) getTimePoint(t int64) *entities.DeviceLogRecord {
	if s.records == nil {
		s.records = make(map[int64]*entities.DeviceLogRecord)
	}
	record := s.records[t]
	if record == nil {
		record = &entities.DeviceLogRecord{
			Time: t,
		}
		s.records[t] = record
	}
	return record
}

func (s *MeasurementsLog) copy(v *entities.DeviceLogRecord) *entities.DeviceLogRecord {
	temperature := *v.Temperature
	return &entities.DeviceLogRecord{
		Time:        v.Time,
		Temperature: &temperature,
	}
}

func (s *MeasurementsLog) Append(v *entities.DeviceLogRecord) {
	point := s.getTimePoint(v.Time)
	point.Temperature = v.Temperature
	// Add others metrics here
}

func (s *MeasurementsLog) ToSlice() []*entities.DeviceLogRecord {
	keys := make([]int64, 0)
	for k := range s.records {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	result := make([]*entities.DeviceLogRecord, 0)
	for _, k := range keys {
		v := s.getTimePoint(k)
		result = append(result, s.copy(v))
	}

	return result
}
