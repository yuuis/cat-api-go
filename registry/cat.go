package registry

import "github.com/yuuis/cat-api-go/usecase"

func (r *registry) NewCatUseCase() usecase.Cat {
	return usecase.NewCat
}
