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
	hash   *hashids.HashID
	logger *log.Logger
}

func NewRegistry(db *gorm.DB, hash *hashids.HashID, logger *log.Logger) Registry {
	return &registry{
		db:     db,
		hash:   hash,
		logger: logger,
	}
}

func (r *registry) NewController() api.Controller {
	return api.NewController(
		usecase.NewInteractor(
			api.NewPresenter(r.logger),
			r.NewCatUseCase(),
		))
}
