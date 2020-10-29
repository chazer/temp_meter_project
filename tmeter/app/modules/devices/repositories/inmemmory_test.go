package repositories

import "testing"

func TestInmemoryRepository(t *testing.T) {
	RepositoryImplementationTestCase(t, struct {
		NewRepo func() DevicesRepositoryInterface
	}{
		NewRepo: func() DevicesRepositoryInterface {
			return MakeDevicesInmemoryRepository()
		},
	})
}
