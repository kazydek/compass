package healthcheck

import (
	"github.com/kyma-incubator/compass/components/director/internal2/model"
)

type inMemoryRepository struct {
	store map[string]*model.Runtime
}

func NewRepository() *inMemoryRepository {
	return &inMemoryRepository{store: make(map[string]*model.Runtime)}
}
