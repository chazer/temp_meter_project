package repositories

import (
	"tmeter/app/modules/users/entities"
)

type UsersRepositoryInterface interface {
	Insert(d *entities.User) *entities.User
	Remove(d *entities.User) bool
	RemoveByUUID(uuid string) bool
	FindByUUID(uuid string) *entities.User
	// TODO: Return Cursor here
	FindByEmail(email string) []*entities.User
	Count() int
	Items() []*entities.User
}
