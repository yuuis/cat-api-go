package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuis/cat-api-go/external/api/utilities"
	"github.com/yuuis/cat-api-go/usecase/cat"
)

func (ctr *controller) GetAllCats(c *gin.Context) {
	ctx := c.Request.Context()
	ctx = utilities.AddGinContext(ctx, c)

	ctr.it.GetAllCats(ctx)
}

func (ctr *controller) GetCat(c *gin.Context) {
	var ipt usecaseCat.GetCatParam
	ipt.ID = c.Param("catID")

	ctx := c.Request.Context()
	ctx = utilities.AddGinContext(ctx, c)

	ctr.it.GetCat(ctx, &ipt)
}

func (ctr *controller) CreateCat(c *gin.Context) {
	var ipt usecaseCat.CreateCatParam
	_ = c.BindJSON(&ipt)

	ctx := c.Request.Context()
	ctx = utilities.AddGinContext(ctx, c)

	ctr.it.CreateCat(ctx, &ipt)
}

func (ctr *controller) UpdateCat(c *gin.Context) {
	var ipt usecaseCat.UpdateCatParam
	_ = c.BindJSON(&ipt)

	ctx := c.Request.Context()
	ctx = utilities.AddGinContext(ctx, c)

	ctr.it.UpdateCat(ctx, &ipt)
}

func (ctr *controller) DeleteCat(c *gin.Context) {
	var ipt usecaseCat.DeleteCatParam
	ipt.ID = c.Param("catID")

	ctx := c.Request.Context()
	ctx = utilities.AddGinContext(ctx, c)

	ctr.it.DeleteCat(ctx, &ipt)
}
