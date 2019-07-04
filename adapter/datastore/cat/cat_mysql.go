package mysql

import (
	"github.com/yuuis/cat-api-go/domain/cat"
)

type Cat struct {
	ID   string
	Name string
}

func FromEntity(e *cat.Cat) *Cat {
	return &Cat {
		ID:   e.ID,
		Name: e.Name,
	}
}

func (c *Cat) ToEntity() *cat.Cat {
	return &cat.Cat {
		ID:   c.ID,
		Name: c.Name,
	}
}
