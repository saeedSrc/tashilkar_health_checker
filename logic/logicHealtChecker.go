package logic

import (
	"fmt"
	"tashilkar_health_checker/domain"
	"tashilkar_health_checker/repo"
)

type HealthChecker interface {
	CreateNewEndPoint(request domain.RegisterApiReq) error
	GetApiLists() (lists []domain.Api, err error)
}

type healthChecker struct {
}

func NewHealthCheckerLogic() HealthChecker {
	h := &healthChecker{}
	return h
}

func (h *healthChecker) CreateNewEndPoint(request domain.RegisterApiReq) error {
	healthRepo := repo.NewHealthCheckerRepo()

	err := healthRepo.InsertNewEndPoint(request)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (h *healthChecker) GetApiLists() ([]domain.Api, error) {
	healthRepo := repo.NewHealthCheckerRepo()
	lists, err := healthRepo.GetApiLists()
	if err != nil {
		return nil, err
	}
	return lists, nil
}
