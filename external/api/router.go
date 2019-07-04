package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuis/cat-api-go/external/api/controllers"
)

func Router(r *gin.Engine, c controllers.Controller) {
	r.GET("/cats", c.GetAllCats)
	r.GET("/cats/:catID", c.GetCat)
	r.POST("/cats", c.CreateCat)
	r.PUT("/cats", c.UpdateCat)
}
