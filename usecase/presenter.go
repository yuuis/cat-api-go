package usecase

import (
	"context"

	"github.com/yuuis/cat-api-go/usecase/output"
)

type Presenter interface {
	ViewCat(ctx context.Context, item *output.Cat)
	ViewError(ctx context.Context, err error)
}
