package input

import "errors"

type GetCat struct {
	ID string `json:"id"`
}

type PostCat struct {
	Name string `json:"name"`
}

func (p *PostCat) ValidateName() error {
	if len(p.Name) < 1 || 25 < len(p.Name) {
		return errors.New("cat name length should be 2 ~ 24")
	}
	return nil
}
