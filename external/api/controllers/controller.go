package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuis/cat-api-go/usecase"
)

type Controller interface {
	GetAllCats(c *gin.Context)
	GetCat(c *gin.Context)
	CreateCat(c *gin.Context)
	UpdateCat(c *gin.Context)
}

type controller struct {
	it usecase.Interactor
}

func NewController(it usecase.Interactor) Controller {
	return &controller{
		it: it,
	}
}
