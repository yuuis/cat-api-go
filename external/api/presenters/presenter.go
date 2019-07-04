package presenters

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/yuuis/cat-api-go/external/api/utilities"
	"github.com/yuuis/cat-api-go/usecase"
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

func (p *presenter) ViewError(ctx context.Context, err error) {
	defer utilities.DeleteGinContext(ctx)
	c := utilities.GetGinContext(ctx)
	p.logger.Println(err)
	p.JSON(c, http.StatusInternalServerError,  map[string]interface{}{"errors":err})
}

func (p *presenter) JSON(c *gin.Context, code int, v interface{}) {
	c.JSON(code, v)
}
