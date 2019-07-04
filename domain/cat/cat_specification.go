package domainCat

import "errors"

func ValidateName(n string) error {
	if len(n) < 1 || 25 < len(n) {
		return errors.New("cat name length should be 2 ~ 24")
	}
	return nil
}
