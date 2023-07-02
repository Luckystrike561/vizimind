package postgres

import (
	"context"
	"fmt"

	"github.com/luckystrike561/vizimind/core/internal/model"
)

func (p *Postgres) UpsertActivity(ctx context.Context, activity *model.Activity) (*model.Activity, error) {
	if err := p.db.WithContext(ctx).Save(activity).Error; err != nil {
		return nil, err
	}

	return activity, nil
}

func (p *Postgres) GetActivity(ctx context.Context, id string) (*model.Activity, error) {
	activity := &model.Activity{}
	if err := p.db.WithContext(ctx).Where("id = ?", id).First(activity).Error; err != nil {
		return nil, err
	}

	return activity, nil
}

func (p *Postgres) ListActivity(ctx context.Context, offset, limit int32) ([]*model.Activity, error) {
	activity := []*model.Activity{}
	if err := p.db.WithContext(ctx).Offset(int(offset)).Limit(int(limit)).Find(&activity).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return activity, nil
}

func (p *Postgres) DeleteActivity(ctx context.Context, id string) error {
	if err := p.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Activity{}).Error; err != nil {
		return err
	}

	return nil
}
