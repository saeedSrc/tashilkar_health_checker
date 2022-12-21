package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"go.uber.org/zap"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoConn *mongo.Client

// This is a user defined method that returns mongo.Client,
// context.Context, context.CancelFunc and error.
// mongo.Client will be used for further database operation.
// context.Context will be used set deadlines for process.
// context.CancelFunc will be used to cancel context and
// resource associated with it.

func connect(uri string) (*mongo.Client, context.Context,
	context.CancelFunc, error) {

	// ctx will be used to set deadline for process, here
	// deadline will of 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

// This is a user defined method that accepts
// mongo.Client and context.Context
// This method used to ping the mongoDB, return error if any.
func ping(client *mongo.Client, ctx context.Context) error {

	// mongo.Client has Ping to ping mongoDB, deadline of
	// the Ping method will be determined by cxt
	// Ping method return error if any occurred, then
	// the error can be handled.
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}

	return nil
}

func createIndexes(client *mongo.Client) {
	indexModels := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{"Url", bsonx.Int32(1)}},
		},
	}

	_, err := client.Database("tashilkar").Collection("healthchecker").Indexes().CreateMany(context.Background(), indexModels)
	if err != nil {
		panic(err)
	}
	_, err = client.Database("tashilkar").Collection("checked_endpoints").Indexes().CreateMany(context.Background(), indexModels)
	if err != nil {
		panic(err)
	}
}

func MongoDBSelection() *mongo.Database {
	return MongoConn.Database("tashilkar")
}

func Init(l *zap.SugaredLogger) *mongo.Client {
	// Get Client, Context, CancelFunc and
	// err from connect method.
	var err error
	var ctx context.Context
	//var cancel context.CancelFunc
	MongoConn, ctx, _, err = connect("mongodb://tashilkar:123456@localhost:27017")
	if err != nil {
		panic(err)
	}

	// one time
	createIndexes(MongoConn)

	// Release resource when the main
	// function is returned.

	// Ping mongoDB with Ping method
	err = ping(MongoConn, ctx)
	if err != nil {
		panic(err)
	}
	l.Info("mongo db connected")
	return MongoConn
}
