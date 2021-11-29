package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *Mongo) InsertOne(database, collection string, options *options.InsertOneOptions, insert interface{}) error {
	coll := m.Database(database).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := coll.InsertOne(ctx, insert, options)
	if err != nil {
		// add more error handler here
		return err
	}
	return err
}

func (m *Mongo) InsertMany(database, collection string, options *options.InsertManyOptions, insert interface{}) error {
	coll := m.Database(database).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := coll.InsertMany(ctx, insert.([]interface{}), options)
	if err != nil {
		// add more error handler here too
		return err
	}
	return nil
}
