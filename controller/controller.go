package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tashilkar_health_checker/domain"
	"tashilkar_health_checker/logic"
)

type Controller interface {
	RegisterNewApi(w http.ResponseWriter, r *http.Request)
	ApiLists(w http.ResponseWriter, r *http.Request)
}

type controller struct {
}

func NewController() Controller {
	c := &controller{}
	return c
}

func (c *controller) RegisterNewApi(w http.ResponseWriter, r *http.Request) {
	request := domain.RegisterApiReq{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fmt.Println(err)
	}

	healthLogic := logic.NewHealthCheckerLogic()

	err = healthLogic.CreateNewEndPoint(request)
	if err != nil {
		fmt.Println(err)
	}

}

func (c *controller) ApiLists(w http.ResponseWriter, r *http.Request) {

	healthLogic := logic.NewHealthCheckerLogic()

	lists, err := healthLogic.GetApiLists()
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
