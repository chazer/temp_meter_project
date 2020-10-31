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

			if i.UUID != f.UUID || i.OwnerEmail != f.OwnerEmail {
				t.Fatalf("Expected the same values")
			}
		})

		t.Run("should keep value copy", func(t *testing.T) {
			a := entities.MakeDevice()
			a.OwnerEmail = "a@local"
			aUUID := insertAndReturnUUID(t, repo, &a)

			// reuse variable
			a.OwnerEmail = "b@local"

			f := findByUUID(t, repo, aUUID)
			if f.OwnerEmail != "a@local" {
				t.Fatalf("Expected previous email")
			}
		})

		t.Run("should not replace item", func(t *testing.T) {
			a := entities.MakeDevice()
			a.OwnerEmail = "a@local"
			aUUID := insertAndReturnUUID(t, repo, &a)

			b := entities.MakeDevice()
			b.UUID = aUUID
			b.OwnerEmail = "b@local"
			repo.Insert(&b)

			f := findByUUID(t, repo, aUUID)
			if f.OwnerEmail != "a@local" {
				t.Fatalf("Expected previous value")
			}
		})

		d1 := entities.MakeDevice()
		d1.OwnerEmail = "d@local"
		d1UUID := insertAndReturnUUID(t, repo, &d1)

		d2 := entities.MakeDevice()
		d2.OwnerEmail = "d@local"
		d2UUID := insertAndReturnUUID(t, repo, &d2)

		t.Run("should return different id for same email", func(t *testing.T) {
			if d1UUID == d2UUID {
				t.Fatalf("UUID should be different")
			}
		})
	})

	t.Run("Remove|RemoveByUUID method", func(t *testing.T) {
		repo := factory.NewRepo()

		a := entities.MakeDevice()
		a.OwnerEmail = "a@local"
		aUUID := insertAndReturnUUID(t, repo, &a)

		b := entities.MakeDevice()
		b.OwnerEmail = "b@local"
		bUUID := insertAndReturnUUID(t, repo, &b)

		c := entities.MakeDevice()
		c.OwnerEmail = "c@local"
		cUUID := insertAndReturnUUID(t, repo, &c)

		f := findByUUID(t, repo, bUUID)
		repo.Remove(f)
		repo.RemoveByUUID(cUUID)

		d1 := entities.MakeDevice()
		d1.OwnerEmail = "d@local"
		d1UUID := insertAndReturnUUID(t, repo, &d1)

		d2 := entities.MakeDevice()
		d2.OwnerEmail = "d@local"
		d2UUID := insertAndReturnUUID(t, repo, &d2)

		repo.RemoveByUUID(d1UUID)

		e := entities.MakeDevice()
		e.OwnerEmail = "e@local"
		eUUID := insertAndReturnUUID(t, repo, &e)
		repo.RemoveByUUID(eUUID)
		insertAndReturnUUID(t, repo, &e)

		t.Run("should be different id", func(t *testing.T) {
			if d1UUID == d2UUID {
				t.Fatalf("UUID should be different")
			}
		})

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

		t.Run("should find item by email after remove", func(t *testing.T) {
			if l := repo.FindByEmail("d@local"); len(l) != 1 {
				t.Fatalf("Expected non empty list")
			}
		})

		t.Run("should find item by email after insert after remove", func(t *testing.T) {
			if l := repo.FindByEmail("e@local"); len(l) != 1 {
				t.Fatalf("Expected non empty list")
			}
		})

	})

	t.Run("FindById method", func(t *testing.T) {
		repo := factory.NewRepo()

		a := entities.MakeDevice()
		a.OwnerEmail = "a@local"
		aUUID := insertAndReturnUUID(t, repo, &a)

		t.Run("should return item", func(t *testing.T) {
			f := findByUUID(t, repo, aUUID)
			if f.OwnerEmail != "a@local" {
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
			b.OwnerEmail = "b@local"

			f := repo.FindByUUID(aUUID)
			if f.OwnerEmail != "a@local" {
				t.Fatalf("Expected unchanged value")
			}
		})
	})

	t.Run("FindByEmail method", func(t *testing.T) {
		repo := factory.NewRepo()

		a := entities.MakeDevice()
		a.OwnerEmail = "a@local"
		repo.Insert(&a)

		b := entities.MakeDevice()
		b.OwnerEmail = "a@local"
		repo.Insert(&b)

		c := entities.MakeDevice()
		c.OwnerEmail = "b@local"
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
			l1[0].OwnerEmail = "c@local"

			l2 := repo.FindByEmail("b@local")
			if l2[0].OwnerEmail != "b@local" {
				t.Fatalf("Expected unchanged value")
			}
		})
	})

	t.Run("Replace method", func(t *testing.T) {

		t.Run("should keep same numbers of items after Replace one", func(t *testing.T) {
			repo := factory.NewRepo()

			a := entities.MakeDevice()
			a.OwnerEmail = "a@local"
			aUUID := insertAndReturnUUID(t, repo, &a)

			b := entities.MakeDevice()
			b.OwnerEmail = "b@local"
			repo.Replace(aUUID, &b)

			if repo.Count() != 1 {
				t.Fatalf("Something wrong with count of items")
			}
		})

		t.Run("should replace one item with another", func(t *testing.T) {
			repo := factory.NewRepo()

			a := entities.MakeDevice()
			a.OwnerEmail = "a@local"
			aUUID := insertAndReturnUUID(t, repo, &a)

			b := entities.MakeDevice()
			b.OwnerEmail = "b@local"
			repo.Replace(aUUID, &b)

			if repo.FindByUUID(b.UUID) == nil {
				t.Fatalf("Replace method will change item id")
			}

			if repo.FindByUUID(aUUID) != nil {
				t.Fatalf("Replace method not remove old item")
			}
		})
	})

	t.Run("Update method", func(t *testing.T) {

		t.Run("should keep same numbers of items after Update one", func(t *testing.T) {
			repo := factory.NewRepo()

			a := entities.MakeDevice()
			a.OwnerEmail = "a@local"
			aUUID := insertAndReturnUUID(t, repo, &a)

			b := entities.MakeDevice()
			b.OwnerEmail = "b@local"
			repo.Update(aUUID, &b)

			if repo.Count() != 1 {
				t.Fatalf("Something wrong with count of items")
			}
		})

		t.Run("should replace one item with another", func(t *testing.T) {
			repo := factory.NewRepo()

			a := entities.MakeDevice()
			a.OwnerEmail = "a@local"
			aUUID := insertAndReturnUUID(t, repo, &a)

			b := entities.MakeDevice()
			b.OwnerEmail = "b@local"
			repo.Update(aUUID, &b)

			c := repo.FindByUUID(aUUID)

			if c == nil {
				t.Fatalf("Update method will remove old item")
			}

			if c.OwnerEmail != "b@local" {
				t.Fatalf("Update method will not change item values")
			}
		})
	})
}
