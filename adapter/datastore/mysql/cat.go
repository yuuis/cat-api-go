package mysql

import (
	"github.com/yuuis/cat-api-go/domain/entity"
	"time"

)

type Cat struct {
	ID int64
	Name string
}

func FromEntity(e *entity.Cat) *Cat {
	return &Cat{
		ID: e.ID,
		Name: e.Name,
	}
}

func (c *Cat) ToEntity() *entity.Cat {
	return &entity.Cat{
		ID: c.ID,
		Name: c.Name,
	}
}
