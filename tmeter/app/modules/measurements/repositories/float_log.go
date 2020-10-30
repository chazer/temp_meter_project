package repositories

type FloatLogRecord struct {
	Time  int64
	Value float32
}

type FloatLogInterface interface {
	Set(time int64, value float32)
	ToSlice() []FloatLogRecord
}

type FloatLog struct {
	records map[int64]float32
}

func (s *FloatLog) Set(time int64, value float32) {
	if s.records == nil {
		s.records = make(map[int64]float32)
	}
	s.records[time] = value
}

func (s *FloatLog) ToSlice() []FloatLogRecord {
	result := make([]FloatLogRecord, 0)
	// TODO
	return result
}
