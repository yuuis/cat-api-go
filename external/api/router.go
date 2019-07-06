package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuis/cat-api-go/external/api/controllers"
)

func Router(e *gin.Engine, c controllers.Controller) {
	v1 := e.Group("/api")

	v1.GET("/cats", c.GetAllCats)
	v1.GET("/cats/:catID", c.GetCat)
	v1.POST("/cats", c.CreateCat)
	v1.PUT("/cats", c.UpdateCat)
	v1.DELETE("/cats/:catID", c.DeleteCat)
}
