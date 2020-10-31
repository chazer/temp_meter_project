package repositories

import (
	"tmeter/app/modules/devices/entities"
)

type DevicesRepositoryInterface interface {
	Insert(d *entities.Device) *entities.Device
	Remove(d *entities.Device) bool
	RemoveByUUID(uuid string) bool
	Replace(uuid string, d *entities.Device)
	Update(uuid string, d *entities.Device)
	FindByUUID(uuid string) *entities.Device
	// TODO: Return Cursor here
	FindByEmail(email string) []*entities.Device
	Count() int
	Items() []*entities.Device
}
