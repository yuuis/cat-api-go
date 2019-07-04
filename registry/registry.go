package registry

import (
	"github.com/jinzhu/gorm"
	"github.com/yuuis/cat-api-go/external/api/controllers"
	"github.com/yuuis/cat-api-go/external/api/presenters"
	"github.com/yuuis/cat-api-go/usecase"
	"log"
)

type Registry interface {
	NewController() controllers.Controller
}

type registry struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewRegistry(db *gorm.DB, logger *log.Logger) Registry {
	return &registry{
		db:     db,
		logger: logger,
	}
}

func (r *registry) NewController() controllers.Controller {
	return controllers.NewController(usecase.NewInteractor(presenters.NewPresenter(r.logger), r.NewCatUseCase()))
}
