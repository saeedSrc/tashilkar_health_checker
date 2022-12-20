package logic

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strings"
	"sync"
	"tashilkar_health_checker/domain"
	"tashilkar_health_checker/repo"
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
	repo repo.HealthChecker
}

func NewHealthCheckerLogic(checker repo.HealthChecker) HealthChecker {
	h := &healthChecker{
		repo: checker,
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

func (h *healthChecker) check(domain, method string, interval int) {
	defer wg.Done()
	for true {
		fmt.Println("checking ", domain, method, interval)
		reader := strings.NewReader(`{}`)
		request, err := http.NewRequest(method, domain, reader)
		if err != nil {
			fmt.Println(err)
		}
		client := &http.Client{
			Timeout: 10 * time.Second,
		}
		resp, err := client.Do(request)
		if err != nil {
			fmt.Println(err)
		}
		if resp != nil {
			if resp.StatusCode >= 500 {
				fmt.Println(err)
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}

}
