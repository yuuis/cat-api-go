package usecase

import (
	"context"
	"github.com/yuuis/cat-api-go/usecase/cat"
)

type Interactor struct {
	presenter Presenter
	cat       usecaseCat.CatUsecase
}

func NewInteractor(pre Presenter, cat usecaseCat.CatUsecase) Interactor {
	return Interactor{
		presenter: pre,
		cat:       cat,
	}
}

func (i *Interactor) GetAllCats(ctx context.Context) {
	cats, err := i.cat.GetAllCats()

	if err != nil {
		i.presenter.ViewError(ctx, err)
		return
	}

	i.presenter.ViewCats(ctx, cats)
}

func (i *Interactor) GetCat(ctx context.Context, ipt *usecaseCat.GetCatParam) {
	cat, err := i.cat.GetCat(ipt)

	if err != nil {
		i.presenter.ViewError(ctx, err)
		return
	}

	i.presenter.ViewCat(ctx, cat)
}

func (i *Interactor) CreateCat(ctx context.Context, ipt *usecaseCat.CreateCatParam) {
	cat, err := i.cat.CreateCat(ipt)

	if err != nil {
		i.presenter.ViewError(ctx, err)
		return
	}

	i.presenter.ViewCat(ctx, cat)
}
