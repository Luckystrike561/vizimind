package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/luckystrike561/vizimind/core/internal/model"
)

func boolPtr(b bool) *bool {
	return &b
}

func int64Ptr(i int64) *int64 {
	return &i
}

func (m *Mongo) UpsertActivity(ctx context.Context, activity *model.Activity) (*model.Activity, error) {
	_, err := m.collection.UpdateByID(ctx, activity.ID, bson.M{"$set": activity}, &options.UpdateOptions{
		Upsert: boolPtr(true),
	})
	if err != nil {
		return nil, err
	}

	return activity, nil
}

func (m *Mongo) GetActivity(ctx context.Context, id string) (*model.Activity, error) {
	activity := &model.Activity{}

	err := m.collection.FindOne(ctx, bson.M{"_id": id}).Decode(activity)
	if err != nil {
		return nil, err
	}

	return activity, nil
}

func (m *Mongo) ListActivity(ctx context.Context, offset, limit int32) ([]*model.Activity, error) {
	activities := []*model.Activity{}

	cursor, err := m.collection.Find(ctx, bson.M{}, &options.FindOptions{
		Skip:  int64Ptr(int64(offset)),
		Limit: int64Ptr(int64(limit)),
	})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &activities); err != nil {
		return nil, err
	}

	return activities, nil
}

func (m *Mongo) DeleteActivity(ctx context.Context, id string) error {
	if err := m.collection.FindOneAndDelete(ctx, bson.M{"_id": id}).Err(); err != nil {
		return err
	}

	return nil
}
