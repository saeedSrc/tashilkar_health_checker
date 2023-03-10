package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"tashilkar_health_checker/config"
	"tashilkar_health_checker/domain"
	"time"
)

type HealthChecker interface {
	InsertNewEndPoint(request domain.RegisterApiReq) error
	InsertCheckedEndPoint(request domain.CheckedApi)
	GetApiLists() ([]domain.Api, error)
	DeleteApi(id primitive.ObjectID) error
	SetStatus(availability domain.HealthCheckerAvailability) error
	GetStatus() (int, error)
}

type healthChecker struct {
	mongo  *DB
	logger *zap.SugaredLogger
	config *config.Config
}

func NewHealthCheckerRepo(mongo *DB, l *zap.SugaredLogger, config *config.Config) HealthChecker {
	h := &healthChecker{
		mongo:  mongo,
		logger: l,
		config: config,
	}
	return h
}

func (h *healthChecker) insertOne(ctx context.Context, col string, doc interface{}) (*mongo.InsertOneResult, error) {
	// select database and collection ith Client.Database method
	// and Database.Collection method
	collection := h.mongo.MongoSelection(col)
	// InsertOne accept two argument of type Context
	// and of empty interface
	result, err := collection.InsertOne(ctx, doc)
	return result, err
}

func (h *healthChecker) InsertNewEndPoint(request domain.RegisterApiReq) error {
	_, err := h.insertOne(context.Background(), h.config.Mongo.Collections[0], request)
	return err
}

func (h *healthChecker) GetApiLists() ([]domain.Api, error) {
	collection := h.mongo.MongoSelection(h.config.Mongo.Collections[0])
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	var results []domain.Api
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (h *healthChecker) DeleteApi(id primitive.ObjectID) error {
	collection := h.mongo.MongoSelection(h.config.Mongo.Collections[0])
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}

func (h *healthChecker) InsertCheckedEndPoint(request domain.CheckedApi) {
	h.insertOne(context.Background(), h.config.Mongo.Collections[1], request)
}

func (h *healthChecker) SetStatus(availability domain.HealthCheckerAvailability) error {
	availability.CreatedAt = time.Now().UTC()
	_, err := h.insertOne(context.Background(), h.config.Mongo.Collections[2], availability)
	return err
}

func (h *healthChecker) GetStatus() (int, error) {
	var availability domain.HealthCheckerAvailability
	collection := h.mongo.MongoSelection(h.config.Mongo.Collections[2])
	filter := bson.D{}
	opts := options.FindOne().SetSort(bson.D{{"createdAt", -1}})
	err := collection.FindOne(context.TODO(), filter, opts).Decode(&availability)
	if err != nil {
		return -1, err
	}

	return availability.Status, nil
}
