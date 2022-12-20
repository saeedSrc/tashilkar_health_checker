package controller

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"net/http"
	"tashilkar_health_checker/domain"
	"tashilkar_health_checker/logic"
)

type Controller interface {
	RegisterNewApi(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	ApiLists(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	DeleteApi(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type controller struct {
	logic  logic.HealthChecker
	logger *zap.SugaredLogger
}

func NewController(logic logic.HealthChecker, logger *zap.SugaredLogger) Controller {
	c := &controller{
		logic:  logic,
		logger: logger,
	}
	return c
}

func (c *controller) RegisterNewApi(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	request := domain.RegisterApiReq{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		c.logger.Errorf("error in deccoding request:%v", err)
		c.response(w, false, 500, nil)
	}

	err = c.logic.CreateNewEndPoint(request)
	if err != nil {
		c.logger.Errorf("error in creating new endpoint: %v", err)
		c.response(w, false, 500, nil)
	}

	c.response(w, true, 200, nil)

}

func (c *controller) ApiLists(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	lists, err := c.logic.GetApiLists()
	if err != nil {
		c.logger.Errorf("error in getting api lists: %v", err)
		c.response(w, false, 500, nil)
	}
	c.response(w, true, 200, lists)
}

func (c *controller) DeleteApi(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	objID, err := primitive.ObjectIDFromHex(ps.ByName("id"))
	if err != nil {
		c.logger.Errorf("error in converting id string to object id: %v", err)
		c.response(w, false, 500, nil)
	}
	err = c.logic.DeleteApi(objID)
	if err != nil {
		c.logger.Errorf("error in deleting record: %v", err)
		c.response(w, false, 500, nil)
	}

	c.response(w, true, 200, nil)

}

func (c *controller) response(w http.ResponseWriter, status bool, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     status,
		"statusCode": code,
		"data":       data,
	})
}
