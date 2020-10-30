package entities

type User struct {
	UUID  string
	Name  string
	Email string
}

func (u User) Copy() *User {
	return &User{
		UUID:  u.UUID,
		Name:  u.Name,
		Email: u.Email,
	}
}
