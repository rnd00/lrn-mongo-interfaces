package main

import (
	"fmt"
	"lrn/mongo-interfaces/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	_ = dbInit()
}

func dbInit() db.Mongo {
	mgc, err := db.NewMongoClient("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	err = mgc.ConnectClient()
	if err != nil {
		panic(err)
	}

	err = mgc.PingServer()
	if err != nil {
		panic(err)
	}

	return *mgc
}

func ExampleInsertOne(m db.MongoCRUD) {
	opt := options.InsertOneOptions{}
	insert := bson.D{bson.E{Key: "key", Value: "value"}}
	err := m.InsertOne("dbName", "collectionName", &opt, insert)
	if err != nil {
		fmt.Println("ExampleInsertOne has failed")
	}
}

func ExampleInsertMany(m db.MongoCRUD) {
	opt := options.InsertManyOptions{}
	doc1 := bson.D{bson.E{Key: "key", Value: "value"}}
	doc2 := bson.D{bson.E{Key: "key", Value: "value"}}
	insert := []bson.D{doc1, doc2}

	err := m.InsertMany("dbName", "collectionName", &opt, insert)
	if err != nil {
		fmt.Println("ExampleInsertMany has failed")
	}
}

func ExampleFindOne(m db.MongoCRUD) {
	opt := options.FindOptions{}

	query := bson.D{bson.E{Key: "key", Value: "value"}}
	output := map[string]interface{}{}

	err := m.FindOne("dbName", "collectionName", query, &opt, &output)
	if err != nil {
		fmt.Println("ExampleFindOne has failed")
	}
}

func ExampleFindMany(m db.MongoCRUD) {
	opt := options.FindOptions{}
	opt.SetMaxTime(30 * time.Second)
	opt.SetLimit(0.0)
	opt.SetProjection(bson.D{bson.E{Key: "key", Value: 1}})

	query := bson.D{bson.E{Key: "key", Value: "value"}}
	output := []map[string]interface{}{}

	err := m.FindOne("dbName", "collectionName", query, &opt, &output)
	if err != nil {
		fmt.Println("ExampleFindMany has failed")
	}
}
