package usecaseCat

func (r *repositories) DeleteCat(ipt *DeleteCatParam) error {
	err := r.catRepository.Delete(ipt.ID)

	if err != nil {
		return err
	}

	return nil
}

type DeleteCatParam struct {
	ID string
}
