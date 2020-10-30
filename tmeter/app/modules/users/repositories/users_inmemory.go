package repositories

import (
	"tmeter/app/modules/users/entities"
	"tmeter/lib/helpers"
)

type UsersInmemoryRepository struct {
	byUUID  map[string]*entities.User
	byEmail map[string][]*entities.User
}

func MakeUsersInmemoryRepository() UsersRepositoryInterface {
	registry := &UsersInmemoryRepository{}
	registry.byUUID = make(map[string]*entities.User)
	registry.byEmail = make(map[string][]*entities.User)
	return registry
}

// Create and Insert UUID into
func (r *UsersInmemoryRepository) Create() *entities.User {
	uuid, _ := helpers.GenUUID()
	d := &entities.User{
		UUID: uuid,
	}
	r.Insert(d)
	return d
}

func (r *UsersInmemoryRepository) Insert(d *entities.User) *entities.User {
	c := d.Copy()

	// Repo has own UUID
	uuid, _ := helpers.GenUUID()
	c.UUID = uuid

	if list := r.byEmail[c.Email]; list != nil {
		r.byEmail[c.Email] = append(list, c)
	} else {
		r.byEmail[c.Email] = []*entities.User{c}
	}

	r.byUUID[c.UUID] = c

	return c
}

func (r *UsersInmemoryRepository) RemoveByUUID(uuid string) bool {
	if f := r.FindByUUID(uuid); f != nil {
		delete(r.byEmail, f.Email)
		delete(r.byUUID, f.UUID)
		return true
	}
	return false
}

func (r *UsersInmemoryRepository) Remove(d *entities.User) bool {
	return r.RemoveByUUID(d.UUID)
}

func (r *UsersInmemoryRepository) FindByUUID(uuid string) *entities.User {
	if d := r.byUUID[uuid]; d != nil {
		return d.Copy()
	}
	return nil
}

func (r *UsersInmemoryRepository) FindByEmail(email string) []*entities.User {
	if list := r.byEmail[email]; list != nil {
		result := make([]*entities.User, len(list))
		for i, v := range list {
			result[i] = v.Copy()
		}
		return result
	}
	return []*entities.User{}
}
func (r *UsersInmemoryRepository) Count() int {
	return len(r.byUUID)
}

func (r *UsersInmemoryRepository) Items() []*entities.User {
	result := make([]*entities.User, len(r.byUUID))
	i := 0
	for _, v := range r.byUUID {
		result[i] = v.Copy()
		i++
	}
	return result
}
