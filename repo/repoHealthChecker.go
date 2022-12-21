package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"tashilkar_health_checker/domain"
)

type HealthChecker interface {
	InsertNewEndPoint(request domain.RegisterApiReq) error
	InsertCheckedEndPoint(request domain.CheckedApi)
	GetApiLists() ([]domain.Api, error)
	DeleteApi(id primitive.ObjectID) error
}

type healthChecker struct {
	mongo  *mongo.Client
	logger *zap.SugaredLogger
}

func NewHealthCheckerRepo(mongo *mongo.Client, l *zap.SugaredLogger) HealthChecker {
	h := &healthChecker{
		mongo:  mongo,
		logger: l,
	}
	return h
}

func (h *healthChecker) insertOne(ctx context.Context, col string, doc interface{}) (*mongo.InsertOneResult, error) {
	// select database and collection ith Client.Database method
	// and Database.Collection method
	db := MongoDBSelection()
	collection := db.Collection(col)

	// InsertOne accept two argument of type Context
	// and of empty interface
	result, err := collection.InsertOne(ctx, doc)
	return result, err
}

func (h *healthChecker) InsertNewEndPoint(request domain.RegisterApiReq) error {
	_, err := h.insertOne(context.Background(), "healthchecker", request)
	return err
}

func (h *healthChecker) GetApiLists() ([]domain.Api, error) {
	db := MongoDBSelection()
	collection := db.Collection("healthchecker")

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
	db := MongoDBSelection()
	collection := db.Collection("healthchecker")

	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}

func (h *healthChecker) InsertCheckedEndPoint(request domain.CheckedApi) {
	h.insertOne(context.Background(), "checked_endpoints", request)
}
