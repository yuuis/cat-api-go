package usecase

import (
	"context"
	"github.com/yuuis/cat-api-go/usecase/input"
)

type Interactor struct {
	presenter Presenter
	cat       Cat
}

func NewInteractor(pre Presenter, cat Cat) Interactor {
	return Interactor{
		presenter: pre,
		cat:       cat,
	}
}

func (i *Interactor) GetCat(ctx context.Context, ipt *input.GetCat) {
	cat, err := i.cat.Get(ipt)

	if err != nil {
		i.presenter.ViewError(ctx, err)
		return
	}

	i.presenter.ViewCat(ctx, cat)
}

func (i *Interactor) PostCat(ctx context.Context, ipt *input.PostCat) {
	cat, err := i.cat.Post(ipt)

	if err != nil {
		i.presenter.ViewError(ctx, err)
		return
	}

	i.presenter.ViewCat(ctx, cat)
}
