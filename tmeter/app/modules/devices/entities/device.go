package entities

import "tmeter/lib/helpers"

type Device struct {
	UUID       string
	Name       string
	OwnerEmail string
}

func MakeDevice() Device {
	uuid, _ := helpers.GenUUID()
	return Device{
		UUID: uuid,
	}
}

func (d Device) Copy() *Device {
	return &Device{
		UUID:       d.UUID,
		Name:       d.Name,
		OwnerEmail: d.OwnerEmail,
	}
}
