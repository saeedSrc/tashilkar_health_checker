package router

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	controller "tashilkar_health_checker/controller"
)
import (
	"net/http"
)

const (
	// PORT local
	PORT = "3001"
)

func RegisterRoutes() {
	fmt.Println("registering routers")
	c := controller.NewController()
	router := httprouter.New()
	router.POST("/api/v1/register", c.RegisterNewApi)
	router.GET("/api/v1/list", c.ApiLists)
	router.GET("/api/v1/delete/:id", c.DeleteApi)
	http.ListenAndServe(":"+PORT, router)

}
