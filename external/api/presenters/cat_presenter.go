package presenters

import (
	"context"
	"github.com/yuuis/cat-api-go/external/api/utilities"
	"github.com/yuuis/cat-api-go/usecase/cat"
	"net/http"
)

func (p *presenter) ViewCat(ctx context.Context, cat *usecaseCat.CatOutput) {
	defer utilities.DeleteGinContext(ctx)
	c := utilities.GetGinContext(ctx)
	p.JSON(c, http.StatusOK, cat)
}

func (p *presenter) ViewCats(ctx context.Context, cats []*usecaseCat.CatOutput) {
	defer utilities.DeleteGinContext(ctx)
	c := utilities.GetGinContext(ctx)
	p.JSON(c, http.StatusOK, cats)
}
