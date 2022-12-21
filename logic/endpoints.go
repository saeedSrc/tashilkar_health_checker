package logic

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"tashilkar_health_checker/config"
	"tashilkar_health_checker/domain"
	"tashilkar_health_checker/repo"
)

type EndPoint interface {
	Create(request domain.RegisterApiReq) error
	GetApiLists() (lists []domain.Api, err error)
	Delete(id primitive.ObjectID) error
}

type endPoint struct {
	repo   repo.HealthChecker
	config *config.Config
}

func NewEndPoint(checker repo.HealthChecker, config *config.Config) EndPoint {
	h := &endPoint{
		repo:   checker,
		config: config,
	}
	return h
}

func (h *endPoint) Create(request domain.RegisterApiReq) error {
	err := h.repo.InsertNewEndPoint(request)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (h *endPoint) GetApiLists() ([]domain.Api, error) {
	lists, err := h.repo.GetApiLists()
	if err != nil {
		return nil, err
	}
	return lists, nil
}

func (h *endPoint) Delete(id primitive.ObjectID) error {
	err := h.repo.DeleteApi(id)
	return err
}
