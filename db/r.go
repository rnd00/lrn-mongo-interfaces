package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindMany uses an abstraction of the find function with a cursor
func (m *Mongo) FindMany(database, collection string, selector bson.D, opts *options.FindOptions, output interface{}) error {
	coll := m.Database(database).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cur, err := coll.Find(ctx, selector, opts)
	if err == mongo.ErrNoDocuments {
		return err
	} else if err != nil {
		return err
	}

	err = cur.All(ctx, output)
	if err != nil {
		return err
	}
	return nil
}

// FindOne
func (m *Mongo) FindOne(database, collection string, selector bson.D, opts *options.FindOneOptions, output interface{}) error {
	coll := m.Database(database).Collection(collection)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := coll.FindOne(ctx, selector).Decode(output)
	if err == mongo.ErrNoDocuments {
		return err
	} else if err != nil {
		return err
	}

	return err
}
