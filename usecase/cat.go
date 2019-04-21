package usecase

import (
	"github.com/yuuis/cat-api-go/domain/entity"
	"github.com/yuuis/cat-api-go/domain/repository"
	"github.com/yuuis/cat-api-go/usecase/input"
	"github.com/yuuis/cat-api-go/usecase/output"
)

type Cat interface {
	Get(ipt *input.GetCat) (*output.Cat, error)
	Post(ipt *input.PostCat) (*output.Cat, error)
}

type cat struct {
	catRepository  repository.Cat
}

func NewCat(catRepository repository.Cat) Cat {
	return &cat{
		catRepository:  catRepository,
	}
}

func (c *cat) Get(ipt *input.GetCat) (*output.Cat, error) {
	cat, err := c.catRepository.Find(ipt.ID)

	if err != nil {
		return nil, err
	}

	oCat := &output.Cat{
		ID:   cat.ID,
		Name: cat.Name,
	}

	return oCat, nil
}

func (c *cat) Post(ipt *input.PostCat) (*output.Cat, error) {
	err := ipt.ValidateName()

	if err != nil {
		return nil, err
	}

	cat, err := c.catRepository.Store(&entity.Cat{
		ID:   entity.GenerateUUID(),
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
