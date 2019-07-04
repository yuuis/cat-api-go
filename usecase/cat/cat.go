package usecaseCat

import (
	"github.com/yuuis/cat-api-go/domain/cat"
)

// TODO: このファイルにあるやつらどこにおけばいいんや
type CatUsecase interface {
	GetAllCats() ([]*CatOutput, error)
	GetCat(ipt *GetCatParam) (*CatOutput, error)
	CreateCat(ipt *CreateCatParam) (*CatOutput, error)
	UpdateCat(ipt *UpdateCatParam) (*CatOutput, error)
	DeleteCat(ipt *DeleteCatParam) error
}

type repositories struct {
	catRepository domainCat.CatRepository
}

func NewCatUsecase(catRepository domainCat.CatRepository) CatUsecase {
	return &repositories{
		catRepository: catRepository,
	}
}

// TODO: outputだけここにあるのきもい
type CatOutput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
