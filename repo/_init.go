package repo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"sync"
	"time"
)

//MongoDB mongodb database connection
var MongoDB *mongo.Client

//once for init mongodb connection
var (
	ctx  context.Context
	once sync.Once
)

//Init manual initialize
func Init() {
	// ensures that our type only gets initialized exactly once.
	once.Do(func() {
		mongoConnection()
	})

}

func mongoConnection() {
	var err error

	var hosts []string

	hosts = append(hosts, "localhost:27017")
	//hosts[0] = "localhost:27017"
MONGOTRY:
	MongoDB, err = mongo.NewClient(
		options.Client().SetHosts(hosts),
		options.Client().SetReadPreference(readpref.PrimaryPreferred()),
		options.Client().SetReplicaSet("tashilkar"),
	)
	if err != nil {
		fmt.Println(err)
		goto MONGOTRY
	}

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

	err = MongoDB.Connect(ctx)
	if err != nil {
		fmt.Println("error 1")
		fmt.Println(err)
	}
	pingContext, can := context.WithTimeout(context.Background(), 10*time.Second)
	defer can()
	err = MongoDB.Ping(pingContext, readpref.PrimaryPreferred())
	if err != nil {
		fmt.Println("error 2")
		fmt.Println(err)
	}
	CreateIndexes()
}

func CreateIndexes() {
	indexModels := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{"Url", bsonx.Int32(1)}},
		},
	}

	_, err := MongoDB.Database("tashilkar").Collection("healthchecker").Indexes().CreateMany(context.Background(), indexModels)
	if err != nil {
		fmt.Println("error 3")
		fmt.Println(err)
	}
}

//MongoDBSelection select and copy mongodb connection
func MongoDBSelection() *mongo.Database {
	return MongoDB.Database("tashilkar")
}

//MongoDBClose close
func MongoDBClose() {
	MongoDB.Disconnect(context.Background())
}
