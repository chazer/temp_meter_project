package repositories

import "testing"

func TestInmemoryRepository(t *testing.T) {
	RepositoryImplementationTestCase(t, struct {
		NewRepo func() UsersRepositoryInterface
	}{
		NewRepo: func() UsersRepositoryInterface {
			return MakeUsersInmemoryRepository()
		},
	})
}
