package handlers

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoHandler struct {
	client *mongo.Client
	database string
	collection string
}

func NewMongoHandler(conn_uri string, database string, collection string) (MongoHandler, time.Duration, error) {
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(conn_uri))
	err = client.Ping(ctx, readpref.Primary())

	defer func() {
		if err != nil {
			panic(err)
		}
	}()

	mhdl := MongoHandler{client, database, collection}

	t := time.Now()
	elapsed := t.Sub(start)

	return mhdl, elapsed, err
}

func (mhdl MongoHandler) ReadAllUsers() (time.Duration, error) {
	start := time.Now()

	collection := mhdl.client.Database(mhdl.database).Collection(mhdl.collection)
	filter := bson.D{{}}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, filter)
	if err != nil {log.Fatalf("[DB_READ_ALL] [FIND] %v", err)}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		//fmt.Println(result)
		if err != nil {log.Fatalf("[DB_READ_ALL] [DECODE] %v", err)}	
	}

	if err := cur.Err(); err != nil {
		log.Fatalf("[DB_READ_ALL] [CURSOR] %v", err)
	}

	t := time.Now()
	elapsed := t.Sub(start)
	return elapsed, err
}

func (mhdl MongoHandler) MakeNewUsers(count int) (time.Duration, time.Duration, error) {
	start := time.Now()

	collection := mhdl.client.Database(mhdl.database).Collection(mhdl.collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	genStart := time.Now()
	users := GenerateRandomUsersAsBSON(count)
	genTime := time.Now().Sub(genStart)

	_, err := collection.InsertMany(ctx, users)
	if err != nil {log.Fatalf("[DB_WRITE] [INSERT] %v", err)}

	t := time.Now()
	elapsed := t.Sub(start)
	return genTime, elapsed, err
}