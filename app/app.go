package app

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"log"
	"tashilkar_health_checker/logic"
	"tashilkar_health_checker/repo"
	"tashilkar_health_checker/router"
)

type App struct {
	Logger    *zap.SugaredLogger
	MongoConn *mongo.Client
	Repo      repo.HealthChecker
	Logic     logic.HealthChecker
}

func NewApp() *App {
	app := &App{}
	return app
}

func (a *App) Init() {
	a.initialLogger()
	a.initRepo()
	a.initialRepoHealthChecker()
	a.initialLogicHealthChecker()
	a.initRouter()
}

func (a *App) initialLogger() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	a.Logger = logger.Sugar()
}

func (a *App) initRepo() {
	a.MongoConn = repo.Init(a.Logger)
}

func (a *App) initRouter() {
	router.RegisterRoutes(a.Logger, a.Logic)
}

func (a *App) initialRepoHealthChecker() {
	a.Repo = repo.NewHealthCheckerRepo(a.MongoConn, a.Logger)
}

func (a *App) initialLogicHealthChecker() {
	a.Logic = logic.NewHealthCheckerLogic(a.Repo)
}
