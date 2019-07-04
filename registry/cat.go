package registry

import (
	"github.com/yuuis/cat-api-go/adapter/datastore/cat"
	"github.com/yuuis/cat-api-go/domain/cat"
	"github.com/yuuis/cat-api-go/usecase/cat"
)

func (r *registry) NewCatUseCase() usecaseCat.CatUsecase {
	return usecaseCat.NewCatUsecase(r.NewCatRepository())
}

func (r *registry) NewCatRepository() domainCat.CatRepository {
	return datastoreCat.NewCatRepository(r.db)
}
