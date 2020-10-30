package repositories

import (
	"tmeter/app/modules/devices/entities"
	"tmeter/lib/helpers"
)

type DevicesInmemoryRepository struct {
	byUUID  map[string]*entities.Device
	byEmail map[string][]*entities.Device
}

func MakeDevicesInmemoryRepository() *DevicesInmemoryRepository {
	registry := &DevicesInmemoryRepository{}
	registry.byUUID = make(map[string]*entities.Device)
	registry.byEmail = make(map[string][]*entities.Device)
	return registry
}

// Create and Insert UUID into
func (r *DevicesInmemoryRepository) Create() *entities.Device {
	uuid, _ := helpers.GenUUID()
	d := &entities.Device{
		UUID: uuid,
	}
	r.Insert(d)
	return d
}

func (r *DevicesInmemoryRepository) Insert(d *entities.Device) *entities.Device {
	c := d.Copy()

	// Repo has own UUID
	uuid, _ := helpers.GenUUID()
	c.UUID = uuid

	if list := r.byEmail[c.OwnerEmail]; list != nil {
		r.byEmail[c.OwnerEmail] = append(list, c)
	} else {
		r.byEmail[c.OwnerEmail] = []*entities.Device{c}
	}

	r.byUUID[c.UUID] = c

	return c
}

func (r *DevicesInmemoryRepository) RemoveByUUID(uuid string) bool {
	if f := r.FindByUUID(uuid); f != nil {
		delete(r.byEmail, f.OwnerEmail)
		delete(r.byUUID, f.UUID)
		return true
	}
	return false
}

func (r *DevicesInmemoryRepository) Remove(d *entities.Device) bool {
	return r.RemoveByUUID(d.UUID)
}

func (r *DevicesInmemoryRepository) FindByUUID(uuid string) *entities.Device {
	if d := r.byUUID[uuid]; d != nil {
		return d.Copy()
	}
	return nil
}

func (r *DevicesInmemoryRepository) FindByEmail(email string) []*entities.Device {
	if list := r.byEmail[email]; list != nil {
		result := make([]*entities.Device, len(list))
		for i, v := range list {
			result[i] = v.Copy()
		}
		return result
	}
	return []*entities.Device{}
}
func (r *DevicesInmemoryRepository) Count() int {
	return len(r.byUUID)
}

func (r *DevicesInmemoryRepository) Items() []*entities.Device {
	result := make([]*entities.Device, len(r.byUUID))
	i := 0
	for _, v := range r.byUUID {
		result[i] = v.Copy()
		i++
	}
	return result
}
