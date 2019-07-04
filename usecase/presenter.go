package usecase

import (
	"context"
	"github.com/yuuis/cat-api-go/usecase/cat"
)

type Presenter interface {
	ViewCat(ctx context.Context, item *usecaseCat.CatOutput)
	ViewCats(ctx context.Context, items []*usecaseCat.CatOutput)
	ViewError(ctx context.Context, err error)
}
