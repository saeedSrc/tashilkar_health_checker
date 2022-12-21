package logic

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"sync"
	"tashilkar_health_checker/config"
	"tashilkar_health_checker/domain"
	"tashilkar_health_checker/repo"
	services "tashilkar_health_checker/service"
	"time"
)

type HealthChecker interface {
	Check() error
	CreateNewEndPoint(request domain.RegisterApiReq) error
	GetApiLists() (lists []domain.Api, err error)
	DeleteApi(id primitive.ObjectID) error
}

var wg sync.WaitGroup

type healthChecker struct {
	repo    repo.HealthChecker
	logger  *zap.SugaredLogger
	service *services.Service
	config  *config.Config
}

func NewHealthCheckerLogic(checker repo.HealthChecker, logger *zap.SugaredLogger,
	service *services.Service, config *config.Config) HealthChecker {
	h := &healthChecker{
		repo:    checker,
		logger:  logger,
		service: service,
		config:  config,
	}
	return h
}

func (h *healthChecker) Check() error {
	apiLists, err := h.repo.GetApiLists()
	if err != nil {
		fmt.Println(err)
	}
	wg.Add(len(apiLists))
	for _, api := range apiLists {
		go h.check(api.Url, api.Method, api.TimeIntervalCheck)
	}
	wg.Wait()
	return nil
}

func (h *healthChecker) CreateNewEndPoint(request domain.RegisterApiReq) error {
	err := h.repo.InsertNewEndPoint(request)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (h *healthChecker) GetApiLists() ([]domain.Api, error) {
	lists, err := h.repo.GetApiLists()
	if err != nil {
		return nil, err
	}
	return lists, nil
}

func (h *healthChecker) DeleteApi(id primitive.ObjectID) error {
	err := h.repo.DeleteApi(id)
	return err
}

func (h *healthChecker) check(url, method string, interval int) {
	defer wg.Done()
	for true {
		reader := strings.NewReader(`{}`)
		request, err := http.NewRequest(method, url, reader)
		if err != nil {
			h.logger.Errorf("could not send request to %s. error is: %v", url, err)
		}
		client := &http.Client{
			Timeout: 10 * time.Second,
		}
		resp, err := client.Do(request)
		if err != nil {
			h.service.Alert(h.config.DownMessage)
		}
		if resp != nil {
			if resp.StatusCode >= 500 {
				h.service.Alert(h.config.DownMessage)
			}
		}
		var r = domain.CheckedApi{
			Method:            method,
			TimeIntervalCheck: int64(interval),
			Url:               url,
			CreatedAt:         time.Now().UTC(),
		}
		h.repo.InsertCheckedEndPoint(r)
		time.Sleep(time.Duration(interval) * time.Second)
	}

}
