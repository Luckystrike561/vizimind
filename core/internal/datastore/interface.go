package datastore

import (
	"context"

	"github.com/luckystrike561/vizimind/core/internal/model"
)

type Datastore interface {
	CreateGPX(ctx context.Context, gpx *model.GPX) error
	GetGPX(ctx context.Context, id string) (*model.GPX, error)
}
