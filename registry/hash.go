package registry

import (
	"github.com/yuuis/cat-api-go/adapter/hashid"
	"github.com/yuuis/cat-api-go/domain/repository"
)

func (r *registry) NewHashRepository() repository.Hash {
	return hashid.NewHash(r.hash)
}
