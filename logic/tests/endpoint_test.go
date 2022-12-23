package tests

import (
	"github.com/golang/mock/gomock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"tashilkar_health_checker/domain"
	mock_repo "tashilkar_health_checker/mocks/repo"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	var request domain.RegisterApiReq
	request.Body = "{\"key1\":\"val1\",\"key2\":\"val2\",\"key3\":val3,\"key4\":\"val4\"}"
	request.Url = "https://varzesh3.com"
	request.Method = "GET"
	request.TimeIntervalCheck = 2
	var header domain.Headers
	header.Authorization = "asdadasd"
	header.XAccessToken = "xxx"
	request.Headers = append(request.Headers, header)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_repo.NewMockHealthChecker(ctrl)
	repo.EXPECT().InsertNewEndPoint(request).Return(nil).AnyTimes()
}

func TestGetLists(t *testing.T) {
	var lists []domain.Api
	var list domain.Api
	list.Url = "google"
	list.Method = "GET"
	list.TimeIntervalCheck = 10
	lists = append(lists, list)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_repo.NewMockHealthChecker(ctrl)
	repo.EXPECT().GetApiLists().Return(lists, nil).AnyTimes()
}

func TestDelete(t *testing.T) {
	objID, _ := primitive.ObjectIDFromHex("63a2153f69987567118606a0")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_repo.NewMockHealthChecker(ctrl)
	repo.EXPECT().DeleteApi(objID).Return(nil).AnyTimes()
}

func TestSetStatus(t *testing.T) {
	var request domain.HealthCheckerAvailability
	request.Status = 1
	request.CreatedAt = time.Now().UTC()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_repo.NewMockHealthChecker(ctrl)
	repo.EXPECT().SetStatus(request).Return(nil).AnyTimes()
}
