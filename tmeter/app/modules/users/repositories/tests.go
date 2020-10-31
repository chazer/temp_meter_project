package repositories

import (
	"testing"
)
import "tmeter/app/modules/users/entities"

func insertAndReturnUUID(t *testing.T, repo UsersRepositoryInterface, item *entities.User) string {
	i := repo.Insert(item)
	if i == nil {
		t.Fatalf("Expected item")
	}
	return i.UUID
}

func findByUUID(t *testing.T, repo UsersRepositoryInterface, uuid string) *entities.User {
	f := repo.FindByUUID(uuid)
	if f == nil {
		t.Fatalf("Expected item")
	}
	return f
}

func RepositoryImplementationTestCase(t *testing.T, factory struct {
	NewRepo func() UsersRepositoryInterface
}) {
	t.Run("Insert method", func(t *testing.T) {
		repo := factory.NewRepo()

		t.Run("should insert", func(t *testing.T) {
			a := entities.User{}

			// inserted
			i := repo.Insert(&a)
			if i == nil {
				t.Fatalf("Expected item")
			}

			// found
			f := findByUUID(t, repo, i.UUID)

			if i.UUID != f.UUID || i.Email != f.Email {
				t.Fatalf("Expected the same values")
			}
		})

		t.Run("should keep value copy", func(t *testing.T) {
			a := entities.User{}
			a.Email = "a@local"
			aUUID := insertAndReturnUUID(t, repo, &a)

			// reuse variable
			a.Email = "b@local"

			f := findByUUID(t, repo, aUUID)
			if f.Email != "a@local" {
				t.Fatalf("Expected previous email")
			}
		})

		t.Run("should not replace item", func(t *testing.T) {
			a := entities.User{}
			a.Email = "a@local"
			aUUID := insertAndReturnUUID(t, repo, &a)

			b := entities.User{}
			b.UUID = aUUID
			b.Email = "b@local"
			repo.Insert(&b)

			f := findByUUID(t, repo, aUUID)
			if f.Email != "a@local" {
				t.Fatalf("Expected previous value")
			}
		})

		d1 := entities.User{}
		d1.Email = "d@local"
		d1UUID := insertAndReturnUUID(t, repo, &d1)

		d2 := entities.User{}
		d2.Email = "d@local"
		d2UUID := insertAndReturnUUID(t, repo, &d2)

		t.Run("should return different id for same email", func(t *testing.T) {
			if d1UUID == d2UUID {
				t.Fatalf("UUID should be different")
			}
		})
	})

	t.Run("Remove|RemoveByUUID method", func(t *testing.T) {
		repo := factory.NewRepo()

		a := entities.User{}
		a.Email = "a@local"
		aUUID := insertAndReturnUUID(t, repo, &a)

		b := entities.User{}
		b.Email = "b@local"
		bUUID := insertAndReturnUUID(t, repo, &b)

		c := entities.User{}
		c.Email = "c@local"
		cUUID := insertAndReturnUUID(t, repo, &c)

		f := findByUUID(t, repo, bUUID)
		repo.Remove(f)
		repo.RemoveByUUID(cUUID)

		d1 := entities.User{}
		d1.Email = "d@local"
		d1UUID := insertAndReturnUUID(t, repo, &d1)

		d2 := entities.User{}
		d2.Email = "d@local"
		d2UUID := insertAndReturnUUID(t, repo, &d2)

		repo.RemoveByUUID(d1UUID)

		e := entities.User{}
		e.Email = "e@local"
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

		a := entities.User{}
		a.Email = "a@local"
		aUUID := insertAndReturnUUID(t, repo, &a)

		t.Run("should return item", func(t *testing.T) {
			f := findByUUID(t, repo, aUUID)
			if f.Email != "a@local" {
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
			b.Email = "b@local"

			f := repo.FindByUUID(aUUID)
			if f.Email != "a@local" {
				t.Fatalf("Expected unchanged value")
			}
		})
	})

	t.Run("FindByEmail method", func(t *testing.T) {
		repo := factory.NewRepo()

		a := entities.User{}
		a.UUID = "id1"
		a.Email = "a@local"
		repo.Insert(&a)

		b := entities.User{}
		b.UUID = "id2"
		b.Email = "a@local"
		repo.Insert(&b)

		c := entities.User{}
		c.UUID = "id3"
		c.Email = "b@local"
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
			l1[0].Email = "c@local"

			l2 := repo.FindByEmail("b@local")
			if l2[0].Email != "b@local" {
				t.Fatalf("Expected unchanged value")
			}
		})
	})

	t.Run("Replace method", func(t *testing.T) {

		t.Run("should keep same numbers of items after Replace one", func(t *testing.T) {
			repo := factory.NewRepo()

			a := entities.User{}
			a.Email = "a@local"
			aUUID := insertAndReturnUUID(t, repo, &a)

			b := entities.User{}
			b.Email = "b@local"
			repo.Replace(aUUID, &b)

			if repo.Count() != 1 {
				t.Fatalf("Something wrong with count of items")
			}
		})

		t.Run("should replace one item with another", func(t *testing.T) {
			repo := factory.NewRepo()

			a := entities.User{}
			a.Email = "a@local"
			aUUID := insertAndReturnUUID(t, repo, &a)

			b := entities.User{}
			b.Email = "b@local"
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

			a := entities.User{}
			a.Email = "a@local"
			aUUID := insertAndReturnUUID(t, repo, &a)

			b := entities.User{}
			b.Email = "b@local"
			repo.Update(aUUID, &b)

			if repo.Count() != 1 {
				t.Fatalf("Something wrong with count of items")
			}
		})

		t.Run("should replace one item with another", func(t *testing.T) {
			repo := factory.NewRepo()

			a := entities.User{}
			a.Email = "a@local"
			aUUID := insertAndReturnUUID(t, repo, &a)

			b := entities.User{}
			b.Email = "b@local"
			repo.Update(aUUID, &b)

			c := repo.FindByUUID(aUUID)

			if c == nil {
				t.Fatalf("Update method will remove old item")
			}

			if c.Email != "b@local" {
				t.Fatalf("Update method will not change item values")
			}
		})
	})
}
