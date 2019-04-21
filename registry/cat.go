package registry

import (
	"github.com/yuuis/cat-api-go/adapter/datastore"
	"github.com/yuuis/cat-api-go/domain/repository"
	"github.com/yuuis/cat-api-go/usecase"
)

func (r *registry) NewCatUseCase() usecase.Cat {
	return usecase.NewCat(r.NewCatRepository())
}

func (r *registry) NewCatRepository() repository.Cat {
	return datastore.NewCat(r.db)
}
