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
	a.Logger.Info("logger initiated")
}

func (a *App) initRepo() {
	a.MongoConn = repo.Init(a.Logger)
	a.Logger.Info("repo initiated")
}

func (a *App) initRouter() {
	router.RegisterRoutes(a.Logger, a.Logic)
	a.Logger.Info("router initiated")
}

func (a *App) initialRepoHealthChecker() {
	a.Repo = repo.NewHealthCheckerRepo(a.MongoConn, a.Logger)
	a.Logger.Info("health repo initiated")
}

func (a *App) initialLogicHealthChecker() {
	a.Logic = logic.NewHealthCheckerLogic(a.Repo)
	a.Logger.Info("health logic initiated")
}
