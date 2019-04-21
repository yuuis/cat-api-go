package api

import "github.com/gin-gonic/gin"

func NewRouter(r *gin.Engine, c Controller) {
	r.GET("/cats/:catID", c.GetCat)
	r.POST("/cats", c.PostCat)
}
