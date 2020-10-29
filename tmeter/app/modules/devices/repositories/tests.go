package repositories

import (
	"testing"
)
import "tmeter/app/modules/devices/entities"

func insertAndReturnUUID(t *testing.T, repo DevicesRepositoryInterface, item *entities.Device) string {
	i := repo.Insert(item)
	if i == nil {
		t.Fatalf("Expected item")
	}
	return i.UUID
}

func findByUUID(t *testing.T, repo DevicesRepositoryInterface, uuid string) *entities.Device {
	f := repo.FindByUUID(uuid)
	if f == nil {
		t.Fatalf("Expected item")
	}
	return f
}

func RepositoryImplementationTestCase(t *testing.T, factory struct {
	NewRepo func() DevicesRepositoryInterface
}) {
	t.Run("Insert method", func(t *testing.T) {
		repo := factory.NewRepo()

		t.Run("should insert", func(t *testing.T) {
			a := entities.MakeDevice()

			// inserted
			i := repo.Insert(&a)
			if i == nil {
				t.Fatalf("Expected item")
			}

			// found
			f := findByUUID(t, repo, i.UUID)

			if i.UUID != f.UUID || i.UserEmail != f.UserEmail {
				t.Fatalf("Expected the same values")
			}
		})

		t.Run("should keep value copy", func(t *testing.T) {
			a := entities.MakeDevice()
			a.UserEmail = "a@local"
			aUUID := insertAndReturnUUID(t, repo, &a)

			// reuse variable
			a.UserEmail = "b@local"

			f := findByUUID(t, repo, aUUID)
			if f.UserEmail != "a@local" {
				t.Fatalf("Expected previous email")
			}
		})

		t.Run("should not replace item", func(t *testing.T) {
			a := entities.MakeDevice()
			a.UserEmail = "a@local"
			aUUID := insertAndReturnUUID(t, repo, &a)

			b := entities.MakeDevice()
			b.UUID = aUUID
			b.UserEmail = "b@local"
			repo.Insert(&b)

			f := findByUUID(t, repo, aUUID)
			if f.UserEmail != "a@local" {
				t.Fatalf("Expected previous value")
			}
		})
	})

	t.Run("Remove|RemoveByUUID method", func(t *testing.T) {
		repo := factory.NewRepo()

		a := entities.MakeDevice()
		a.UserEmail = "a@local"
		aUUID := insertAndReturnUUID(t, repo, &a)

		b := entities.MakeDevice()
		b.UserEmail = "b@local"
		bUUID := insertAndReturnUUID(t, repo, &b)

		c := entities.MakeDevice()
		c.UserEmail = "c@local"
		cUUID := insertAndReturnUUID(t, repo, &c)

		f := findByUUID(t, repo, bUUID)
		repo.Remove(f)
		repo.RemoveByUUID(cUUID)

		t.Run("should find other item after remove", func(t *testing.T) {
			findByUUID(t, repo, aUUID)
		})

		t.Run("should not find item by id after remove", func(t *testing.T) {
			if nothing := repo.FindByUUID(bUUID); nothing != nil {
				t.Fatalf("Expected nil")
			}
			if nothing := repo.FindByUUID(cUUID); nothing != nil {
				t.Fatalf("Expected nil")
			}
		})

		t.Run("should not find item by email after remove", func(t *testing.T) {
			if l := repo.FindByEmail("b@local"); len(l) != 0 {
				t.Fatalf("Expected empty list")
			}
			if l := repo.FindByEmail("c@local"); len(l) != 0 {
				t.Fatalf("Expected empty list")
			}
		})

	})

	t.Run("FindById method", func(t *testing.T) {
		repo := factory.NewRepo()

		a := entities.MakeDevice()
		a.UserEmail = "a@local"
		aUUID := insertAndReturnUUID(t, repo, &a)

		t.Run("should return item", func(t *testing.T) {
			f := findByUUID(t, repo, aUUID)
			if f.UserEmail != "a@local" {
				t.Fatalf("Expected item")
			}
		})

		t.Run("should return nil", func(t *testing.T) {
			f := repo.FindByUUID("invalid")
			if f != nil {
				t.Fatalf("Expected nil")
			}
		})

		t.Run("should return copy", func(t *testing.T) {
			b := repo.FindByUUID(aUUID)
			b.UserEmail = "b@local"

			f := repo.FindByUUID(aUUID)
			if f.UserEmail != "a@local" {
				t.Fatalf("Expected unchanged value")
			}
		})
	})

	t.Run("FindByEmail method", func(t *testing.T) {
		repo := factory.NewRepo()

		a := entities.MakeDevice()
		a.UUID = "id1"
		a.UserEmail = "a@local"
		repo.Insert(&a)

		b := entities.MakeDevice()
		b.UUID = "id2"
		b.UserEmail = "a@local"
		repo.Insert(&b)

		c := entities.MakeDevice()
		c.UUID = "id3"
		c.UserEmail = "b@local"
		repo.Insert(&c)

		t.Run("should return two", func(t *testing.T) {
			l := repo.FindByEmail("a@local")
			if len(l) != 2 {
				t.Fatalf("Expected list with one item")
			}
		})

		t.Run("should return one", func(t *testing.T) {
			l := repo.FindByEmail("b@local")
			if len(l) != 1 {
				t.Fatalf("Expected list with one item")
			}
		})

		t.Run("should return empty list", func(t *testing.T) {
			l := repo.FindByEmail("c@local")
			if len(l) != 0 {
				t.Fatalf("Expected empty list")
			}
		})

		t.Run("should return copy", func(t *testing.T) {
			l1 := repo.FindByEmail("b@local")
			l1[0].UserEmail = "c@local"

			l2 := repo.FindByEmail("b@local")
			if l2[0].UserEmail != "b@local" {
				t.Fatalf("Expected unchanged value")
			}
		})
	})
}
