package repositories

import "testing"

func TestMeasurementsLog(t *testing.T) {
	log := MeasurementsLog{}

	t.Run("getTimePoint method", func(t *testing.T) {
		p := log.getTimePoint(123)

		t.Run("should create struct with same time", func(t *testing.T) {
			if p.Time != 123 {
				t.Error("Time mismatched")
			}
		})

		t.Run("should always return same reference", func(t *testing.T) {
			a := log.getTimePoint(123)
			b := log.getTimePoint(123)
			v := float32(12.34)
			a.Temperature = &v

			if a.Temperature != b.Temperature {
				t.Error("Different references")
			}
		})
	})
}
