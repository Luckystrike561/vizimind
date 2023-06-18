package datastore

import (
	"context"

	"github.com/luckystrike561/vizimind/core/internal/model"
)

type Datastore interface {
	UpsertGPX(ctx context.Context, gpx *model.GPX) (*model.GPX, error)
	GetGPX(ctx context.Context, id string) (*model.GPX, error)
	ListGPX(ctx context.Context, offset, limit int32) ([]*model.GPX, error)
	DeleteGPX(ctx context.Context, id string) error
}
