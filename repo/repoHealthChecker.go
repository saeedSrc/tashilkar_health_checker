package repo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"tashilkar_health_checker/domain"
)

type HealthChecker interface {
	Check() error
	InsertNewEndPoint(request domain.RegisterApiReq) error
	GetApiLists() ([]domain.Api, error)
	DeleteApi(id primitive.ObjectID) error
}

type healthChecker struct {
}

func NewHealthCheckerRepo() HealthChecker {
	h := &healthChecker{}
	return h
}

func (h *healthChecker) Check() error {
	apiLists, err := h.GetApiLists()
	if err != nil {
		return err
	}
	for _, api := range apiLists {
		fmt.Println(api.TimeIntervalCheck)
		fmt.Println(api.Url)
		fmt.Println(api.Method)
	}
	return nil
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
	if err != nil {
		fmt.Println("error on inserting ", err)
	}
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
