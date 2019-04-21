package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/yuuis/cat-api-go/usecase"
	"github.com/yuuis/cat-api-go/usecase/output"
	"log"
	"net/http"
)

type presenter struct {
	logger *log.Logger
}

func NewPresenter(logger *log.Logger) usecase.Presenter {
	return &presenter{
		logger: logger,
	}
}

func (p *presenter) ViewCat(ctx context.Context, cat *output.Cat) {
	defer deleteGinContext(ctx)
	c := getGinContext(ctx)
	p.JSON(c, http.StatusOK, cat)
}

func (p *presenter) ViewError(ctx context.Context, err error) {
	defer deleteGinContext(ctx)
	c := getGinContext(ctx)
	p.logger.Println(err)
	p.JSON(c, http.StatusInternalServerError, "")
}

func (p *presenter) JSON(c *gin.Context, code int, v interface{}) {
	c.JSON(code, v)
}
