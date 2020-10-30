package views

import (
	"tmeter/app/modules/devices/entities"
	"tmeter/lib/api/views"
)

type structDeviceData struct {
	UUID  string `json:"uuid"`
	Email string `json:"email"`
}

type structDeviceDoc struct {
	Device structDeviceData `json:"device"`
}

type deviceScheme struct {
	views.StructTaggingScheme
}

func (w *deviceScheme) ToTaggedStruct(i interface{}) (interface{}, error) {
	d := i.(*entities.Device)
	return structDeviceDoc{
		Device: structDeviceData{
			UUID:  d.UUID,
			Email: d.OwnerEmail,
		},
	}, nil
}

func NewDeviceApiView() views.ApiViewInterface {
	return &views.ApiView{
		Scheme: &deviceScheme{},
	}
}
