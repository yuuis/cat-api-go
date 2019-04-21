package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuis/cat-api-go/usecase"
	"github.com/yuuis/cat-api-go/usecase/input"
)

type Controller interface {
	GetCat(c *gin.Context)
	PostCat(c *gin.Context)
}

type controller struct {
	it usecase.Interactor
}

func NewController(it usecase.Interactor) Controller {
	return &controller{
		it: it,
	}
}

func (ctr *controller) GetCat(c *gin.Context) {
	var ipt input.GetCat
	ipt.ID = c.Param("catID")

	ctx := c.Request.Context()
	ctx = addGinContext(ctx, c)

	ctr.it.GetCat(ctx, &ipt)
}

func (ctr *controller) PostCat(c *gin.Context) {
	var ipt input.PostCat
	_ = c.BindJSON(&ipt)

	ctx := c.Request.Context()
	ctx = addGinContext(ctx, c)

	ctr.it.PostCat(ctx, &ipt)
}
