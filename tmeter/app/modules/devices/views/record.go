package views

import (
	"time"
	"tmeter/app/modules/devices/entities"
	"tmeter/lib/api/views"
)

// 	views.StructTaggingScheme
type logRecordScheme struct {
	Timestamp   int64     `json:"ts"`
	Time        time.Time `json:"time"`
	Temperature float32   `json:"temperature"`
}

func (w *logRecordScheme) ToTaggedStruct(i interface{}) (interface{}, error) {
	d := i.(*entities.DeviceLogRecord)
	return logRecordScheme{
		Timestamp:   d.Time,
		Time:        time.Unix(d.Time/1000, d.Time%1000*1000000),
		Temperature: *d.Temperature,
	}, nil
}

func NewDeviceLogPointApiView() views.ApiViewInterface {
	return &views.ApiView{
		Scheme: &logRecordScheme{},
	}
}
