package router

import (
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
	controller "tashilkar_health_checker/controller"
	"tashilkar_health_checker/logic"
)
import (
	"net/http"
)

const (
	// PORT local
	PORT = "3001"
)

func RegisterRoutes(l *zap.SugaredLogger, logic logic.EndPoint) {
	l.Info("registering routers")
	c := controller.NewController(logic, l)
	router := httprouter.New()
	router.POST("/api/v1/register", c.RegisterNewApi)
	router.GET("/api/v1/list", c.ApiLists)
	router.GET("/api/v1/delete/:id", c.DeleteApi)
	http.ListenAndServe(":"+PORT, router)

}
