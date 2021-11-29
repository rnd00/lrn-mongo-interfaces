package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *Mongo) DeleteOne(database, collection string, options *options.DeleteOptions, selector bson.D) error {
	coll := m.Database(database).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := coll.DeleteOne(ctx, selector, options)
	if err != nil {
		return err
	}
	return err
}

func (m *Mongo) DeleteMany(database, collection string, options *options.DeleteOptions, selector bson.D) error {
	coll := m.Database(database).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := coll.DeleteMany(ctx, selector, options)
	if err != nil {
		return err
	}
	return err
}
