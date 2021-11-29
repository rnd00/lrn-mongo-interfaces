package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *Mongo) UpdateOne(database, collection string, selector bson.D, opts *options.UpdateOptions, update interface{}) error {
	coll := m.Database(database).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := coll.UpdateOne(ctx, selector, update, opts)
	if err != nil {
		return err
	}
	return nil
}

func (m *Mongo) UpdateMany(database, collection string, selector bson.D, opts *options.UpdateOptions, update interface{}) error {
	coll := m.Database(database).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := coll.UpdateMany(ctx, selector, update, opts)
	if err != nil {
		return err
	}
	return nil
}
