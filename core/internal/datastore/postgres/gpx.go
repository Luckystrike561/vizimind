package postgres

import (
	"context"

	"github.com/luckystrike561/vizimind/core/internal/model"
)

func (p *Postgres) UpsertGPX(ctx context.Context, gpx *model.GPX) (*model.GPX, error) {
	if err := p.db.WithContext(ctx).Save(gpx).Error; err != nil {
		return nil, err
	}

	return gpx, nil
}

func (p *Postgres) GetGPX(ctx context.Context, id string) (*model.GPX, error) {
	gpx := &model.GPX{}
	if err := p.db.WithContext(ctx).Where("id = ?", id).First(gpx).Error; err != nil {
		return nil, err
	}

	return gpx, nil
}

func (p *Postgres) ListGPX(ctx context.Context, offset, limit int32) ([]*model.GPX, error) {
	gpx := []*model.GPX{}
	if err := p.db.WithContext(ctx).Offset(int(offset)).Limit(int(limit)).Find(&gpx).Error; err != nil {
		return nil, err
	}

	return gpx, nil
}

func (p *Postgres) DeleteGPX(ctx context.Context, id string) error {
	if err := p.db.WithContext(ctx).Where("id = ?", id).Delete(&model.GPX{}).Error; err != nil {
		return err
	}

	return nil
}
