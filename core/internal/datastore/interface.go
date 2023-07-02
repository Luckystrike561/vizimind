package datastore

import (
	"context"

	"github.com/luckystrike561/vizimind/core/internal/model"
)

type Datastore interface {
	UpsertActivity(ctx context.Context, activity *model.Activity) (*model.Activity, error)
	GetActivity(ctx context.Context, id string) (*model.Activity, error)
	ListActivity(ctx context.Context, offset, limit int32) ([]*model.Activity, error)
	DeleteActivity(ctx context.Context, id string) error
}
