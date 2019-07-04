package usecaseCat

import (
	domainCat "github.com/yuuis/cat-api-go/domain/cat"
)

func (r *repositories) UpdateCat(ipt *UpdateCatParam) (*CatOutput, error) {
	err := domainCat.ValidateName(ipt.Name)

	if err != nil {
		return nil, err
	}

	cat, err := r.catRepository.Find(ipt.ID)

	if err != nil {
		return nil, err
	}

	cat, err = r.catRepository.Update(&domainCat.Cat{
		ID:   ipt.ID,
		Name: ipt.Name,
	})

	if err != nil {
		return nil, err
	}

	oCat := &CatOutput{
		ID:   cat.ID,
		Name: cat.Name,
	}

	return oCat, nil

}

type UpdateCatParam struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
