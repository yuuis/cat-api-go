package datastoreCat

import (
	"github.com/yuuis/cat-api-go/domain/cat"
)

type Cat struct {
	ID   string
	Name string
}

func FromEntity(e *domainCat.Cat) *Cat {
	return &Cat{
		ID:   e.ID,
		Name: e.Name,
	}
}

func (c *Cat) ToEntity() *domainCat.Cat {
	return &domainCat.Cat{
		ID:   c.ID,
		Name: c.Name,
	}
}
