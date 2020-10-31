package entities

import (
	"time"
	"tmeter/lib/helpers"
)

type Device struct {
	UUID       string
	Name       string
	OwnerEmail string
	UpdatedAt  time.Time
}

func MakeDevice() Device {
	uuid, _ := helpers.GenUUID()
	return Device{
		UUID:      uuid,
		UpdatedAt: time.Now(),
	}
}

func (d Device) Copy() *Device {
	return &Device{
		UUID:       d.UUID,
		Name:       d.Name,
		OwnerEmail: d.OwnerEmail,
		UpdatedAt:  d.UpdatedAt,
	}
}
