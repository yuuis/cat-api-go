package usecaseCat

func (r *repositories) GetCat(ipt *GetCatParam) (*CatOutput, error) {
	cat, err := r.catRepository.Find(ipt.ID)

	if err != nil {
		return nil, err
	}

	oCat := &CatOutput{
		ID:   cat.ID,
		Name: cat.Name,
	}

	return oCat, nil
}

type GetCatParam struct {
	ID string
}
