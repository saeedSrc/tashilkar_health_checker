package router

import (
	"fmt"
	controller "tashilkar_health_checker/controller"
)
import (
	"log"
	"net/http"
)

const (
	// PORT local
	PORT = "3001"
)

func RegisterRoutes() {
	fmt.Println("registering routers")
	c := controller.NewController()
	http.HandleFunc("/api/v1/fake1", c.RegisterNewApi)
	http.HandleFunc("/api/v1/fake2", c.RegisterNewApi)
	http.HandleFunc("/api/v1/fake3", c.RegisterNewApi)
	http.HandleFunc("/api/v1/register", c.RegisterNewApi)
	http.HandleFunc("/api/v1/list", c.ApiLists)
	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		log.Fatal(err)
	}

}
