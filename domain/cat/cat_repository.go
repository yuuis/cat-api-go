package domainCat

type CatRepository interface {
	All() ([]*Cat, error)
	Find(id string) (*Cat, error)
	Store(*Cat) (*Cat, error)
	Update(*Cat) (*Cat, error)
	Delete(id string) error
}
