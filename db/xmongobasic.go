package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// NewMongoClient is a constructor for a mongo client, it has to connect to outside
func NewMongoClient(mongoConn string) (*Mongo, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoConn))
	if err != nil {
		return nil, err
	}
	mg := Mongo{
		client,
	}
	return &mg, nil
}

// PingServer pings the server
func (m *Mongo) PingServer() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := m.Ping(ctx, readpref.Primary())

	return err
}

// ConnectClient connects the client instance
func (m *Mongo) ConnectClient() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := m.Connect(ctx)

	return err
}
