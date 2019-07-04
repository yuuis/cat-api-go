package registry

import (
	"github.com/jinzhu/gorm"
	"github.com/yuuis/cat-api-go/external/api"
	"github.com/yuuis/cat-api-go/usecase"
	"log"
)

type Registry interface {
	NewController() api.Controller
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

func (r *registry) NewController() api.Controller {
	return api.NewController(usecase.NewInteractor(api.NewPresenter(r.logger), r.NewCatUseCase()))
}
