package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuis/cat-api-go/usecase"
	"github.com/yuuis/cat-api-go/usecase/cat"
)

type Controller interface {
	GetAllCats(c *gin.Context)
	GetCat(c *gin.Context)
	CreateCat(c *gin.Context)
}

type controller struct {
	it usecase.Interactor
}

func NewController(it usecase.Interactor) Controller {
	return &controller {
		it: it,
	}
}

func (ctr *controller) GetAllCats(c *gin.Context) {
	ctx := c.Request.Context()
	ctx = addGinContext(ctx, c)

	ctr.it.GetAllCats(ctx)
}

func (ctr *controller) GetCat(c *gin.Context) {
	var ipt usecaseCat.GetCatParam
	ipt.ID = c.Param("catID")

	ctx := c.Request.Context()
	ctx = addGinContext(ctx, c)

	ctr.it.GetCat(ctx, &ipt)
}

func (ctr *controller) CreateCat(c *gin.Context) {
	var ipt usecaseCat.CreateCatParam
	_ = c.BindJSON(&ipt)

	ctx := c.Request.Context()
	ctx = addGinContext(ctx, c)

	ctr.it.CreateCat(ctx, &ipt)
}
