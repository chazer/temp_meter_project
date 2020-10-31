package repositories

import "tmeter/app/modules/devices/entities"

type MeasurementsLogInterface interface {
	Append(v *entities.DeviceLogRecord)
	ToSlice() []*entities.DeviceLogRecord
}

type DeviceLogInterface = interface {
	MeasurementsLogInterface
}
