package usecaseCat

import (
	"github.com/yuuis/cat-api-go/domain"
	"github.com/yuuis/cat-api-go/domain/cat"
)

func (r *repositories) CreateCat(ipt *CreateCatParam) (*CatOutput, error) {
	err := domainCat.ValidateName(ipt.Name)

	if err != nil {
		return nil, err
	}

	cat, err := r.catRepository.Store(&domainCat.Cat {
		ID:   domain.GenerateUUID(),
		Name: ipt.Name,
	})

	if err != nil {
		return nil, err
	}

	oCat := &CatOutput {
		ID:   cat.ID,
		Name: cat.Name,
	}

	return oCat, nil
}

type CreateCatParam struct {
	Name string `json:"name"`
}
