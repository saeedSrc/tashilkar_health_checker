package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"go.uber.org/zap"
	"tashilkar_health_checker/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DB struct {
	MongoConn *mongo.Client
	logger    *zap.SugaredLogger
	config    *config.Config
}

func NewDB(logger *zap.SugaredLogger, config *config.Config) *DB {
	return &DB{
		logger: logger,
		config: config,
	}
}

func (d *DB) connect(uri string) (*mongo.Client, context.Context,
	context.CancelFunc, error) {

	// ctx will be used to set deadline for process, here
	// deadline will of 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func (d *DB) ping(ctx context.Context) error {
	// mongo.Client has Ping to ping mongoDB, deadline of
	// the Ping method will be determined by cxt
	// Ping method return error if any occurred, then
	// the error can be handled.
	if err := d.MongoConn.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}

	return nil
}

func (d *DB) createIndexes() {
	indexModels := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{"Url", bsonx.Int32(1)}},
		},
	}
	for _, v := range d.config.Mongo.Collections {
		_, err := d.MongoConn.Database(d.config.Mongo.Database).Collection(v).Indexes().CreateMany(context.Background(), indexModels)
		if err != nil {
			d.logger.Errorf("could not create index. err is: %v", err)
			panic(err)
		}
	}
}

func (d *DB) MongoSelection(col string) *mongo.Collection {
	return d.MongoConn.Database(d.config.Mongo.Database).Collection(col)
}

func (d *DB) Init() *DB {
	// Get Client, Context, CancelFunc and
	// err from connect method.
	var err error
	var ctx context.Context
	//var cancel context.CancelFunc
	d.logger.Infof("trying to conect to mongo: %s", d.config.Mongo.Uri)
	d.MongoConn, ctx, _, err = d.connect(d.config.Mongo.Uri)
	if err != nil {
		d.logger.Error("could not connect mongo")
		panic(err)
	}

	// one time
	d.createIndexes()

	// Release resource when the main
	// function is returned.

	// Ping mongoDB with Ping method
	err = d.ping(ctx)
	if err != nil {
		d.logger.Error("could not ping mongo")
		panic(err)
	}
	d.logger.Info("mongo db connected")
	return d
}
