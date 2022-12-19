package main

import (
	"tashilkar_health_checker/repo"
	"tashilkar_health_checker/router"
)

func main() {
	repo.Init()
	router.RegisterRoutes()
}
