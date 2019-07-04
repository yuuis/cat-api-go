package usecase

import (
	"github.com/yuuis/cat-api-go/domain"
	"github.com/yuuis/cat-api-go/domain/cat"
	"github.com/yuuis/cat-api-go/usecase/input"
	"github.com/yuuis/cat-api-go/usecase/output"
)

type CatUsecase interface {
	Get(ipt *input.GetCat) (*output.Cat, error)
	Post(ipt *input.CreateCat) (*output.Cat, error)
}

type catRepository struct {
	catRepository cat.CatRepository
}

func NewCat(catRepository catRepository) CatUsecase {
	return &catRepository {
		catRepository:  catRepository,
	}
}

//func (c *catRepository) Get(ipt *input.GetCat) (*output.Cat, error) {
//	cat, err := c.catRepository.Find(ipt.ID)
//
//	if err != nil {
//		return nil, err
//	}
//
//	oCat := &output.Cat{
//		ID:   cat.ID,
//		Name: cat.Name,
//	}
//
//	return oCat, nil
//}


func (c *catRepository) Post(ipt *input.CreateCat) (*output.Cat, error) {
	err := ipt.ValidateName()

	if err != nil {
		return nil, err
	}

	cat, err := c.catRepository.Store(&cat.Cat {
		ID:   domain.GenerateUUID(),
		Name: ipt.Name,
	})

	if err != nil {
		return nil, err
	}

	oCat := &output.Cat{
		ID:   cat.ID,
		Name: cat.Name,
	}

	return oCat, nil
}
