package postgres

import (
	"context"

	"github.com/luckystrike561/vizimind/core/internal/model"
)

func (p *Postgres) CreateGPX(ctx context.Context, gpx *model.GPX) error {
	return p.db.WithContext(ctx).Create(gpx).Error
}

func (p *Postgres) GetGPX(ctx context.Context, id string) (*model.GPX, error) {
	gpx := &model.GPX{}
	if err := p.db.WithContext(ctx).Where("id = ?", id).First(gpx).Error; err != nil {
		return nil, err
	}

	return gpx, nil
}
