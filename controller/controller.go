package controller

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	logic logic.HealthChecker
}

func NewController(logic logic.HealthChecker) Controller {
	c := &controller{
		logic: logic,
	}
	return c
}

func (c *controller) RegisterNewApi(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	request := domain.RegisterApiReq{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fmt.Println(err)
	}

	err = c.logic.CreateNewEndPoint(request)
	if err != nil {
		fmt.Println(err)
	}

}

func (c *controller) ApiLists(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	lists, err := c.logic.GetApiLists()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("here")
	fmt.Println(len(lists))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     "success",
		"statusCode": 200,
		"data":       lists,
	})

}

func (c *controller) DeleteApi(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	objID, err := primitive.ObjectIDFromHex(ps.ByName("id"))
	if err != nil {
		panic(err)
	}

	fmt.Println(objID)
	err = c.logic.DeleteApi(objID)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     "success",
		"statusCode": 200,
		"data":       nil,
	})

}
