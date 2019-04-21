package repository

import "github.com/yuuis/cat-api-go/domain/entity"

type Cat interface {
	All() ([]*entity.Cat, error)
	Find(id int64) (*entity.Cat, error)
	Store(*entity.Cat) (*entity.Cat, error)
}
