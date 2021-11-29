package db

import (
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo abstract the mongo client to use the functions
type Mongo struct {
	*mongo.Client
}

// MongoCRUD is an interface to abstract and test the packages
type MongoCRUD interface {
	FindMany(database, collection string, selector bson.D, options *options.FindOptions, output interface{}) error
	FindOne(database, collection string, selector bson.D, options *options.FindOptions, output interface{}) error
	InsertOne(database, collection string, options *options.InsertOneOptions, insert interface{}) error
	InsertMany(database, collection string, options *options.InsertManyOptions, insert interface{}) error
	Upsert(database, collection string, selector bson.D, options *options.UpdateOptions, update interface{}) error
	Update(database, collection string, selector bson.D, options *options.UpdateOptions, update interface{}) error
	Delete(database, collection string, options *options.DeleteOptions, selector bson.D) error
}

// MongoBasic will extend the CRUD functionality and adding client functions
type MongoBasic interface {
	PingServer() error
	ConnectClient() error
	MongoCRUD
}

// MongoStruct will help on query making
type MongoStruct struct {
	Field   string    `json:"field,omitempty"`
	Array   []string  `json:"array,omitempty"`
	Integer int       `json:"integer,omitempty"`
	Ptr     *string   `json:"ptr,omitempty"`
	DtRange DateRange `json:"dtRange,omitempty"`
}

// DateRange is a struct that helps with time range checks
type DateRange struct {
	From *time.Time `json:"$gte,omitempty"`
	To   *time.Time `json:"$lte,omitempty"`
}

func BuildBsonQuery(search map[string]interface{}) bson.D {
	terms := []bson.D{}
	query := bson.D{}

	for key, v := range search {
		if reflect.ValueOf(v).Kind() == reflect.Slice {
			in := bson.D{bson.E{Key: "$in", Value: v}}
			terms = append(terms, bson.D{bson.E{Key: key, Value: in}})
			continue
		}
		if v != nil {
			singleTerm := bson.D{bson.E{Key: "$in", Value: v}}
			terms = append(terms, singleTerm)
		}
	}
	query = append(query, bson.E{Key: "$and", Value: terms})
	return query
}
