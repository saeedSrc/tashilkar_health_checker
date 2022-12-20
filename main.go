package main

import (
	"tashilkar_health_checker/app"
	"tashilkar_health_checker/router"
)

func main() {
	application := app.NewApp()
	application.Init()
	router.RegisterRoutes()
}
