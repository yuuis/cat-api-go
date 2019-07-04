package usecaseCat

func (r *repositories) GetAllCats() ([]*CatOutput, error) {
	cats, err := r.catRepository.All()

	if err != nil {
		return nil, err
	}

	var oCats []*CatOutput

	for _, cat := range cats {
		oCats = append(oCats, &CatOutput {
			ID: cat.ID,
			Name: cat.Name,
		})
	}

	return oCats, nil
}
