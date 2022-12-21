package app

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"log"
	"tashilkar_health_checker/config"
	"tashilkar_health_checker/logic"
	"tashilkar_health_checker/repo"
	"tashilkar_health_checker/router"
	services "tashilkar_health_checker/service"
)

type App struct {
	Config        *config.Config
	Logger        *zap.SugaredLogger
	MongoConn     *mongo.Client
	Repo          repo.HealthChecker
	HealthLogic   logic.HealthChecker
	EndPointLogic logic.EndPoint
	Service       *services.Service
}

func NewApp() *App {
	app := &App{}
	return app
}

func (a *App) Init() {
	a.initConfig()
	a.initialLogger()
	a.initService()
	a.initRepo()
	a.initialRepoHealthChecker()
	a.initialLogic()
	a.initRouter()
}
func (a *App) initConfig() {
	a.Config = config.Init("./config.yaml")
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
	router.RegisterRoutes(a.Logger, a.EndPointLogic)
	a.Logger.Info("router initiated")
}

func (a *App) initialRepoHealthChecker() {
	a.Repo = repo.NewHealthCheckerRepo(a.MongoConn, a.Logger)
	a.Logger.Info("health repo initiated")
}

func (a *App) initialLogic() {
	a.HealthLogic = logic.NewHealthCheckerLogic(a.Repo, a.Logger, a.Service, a.Config)
	a.EndPointLogic = logic.NewEndPoint(a.Repo, a.Config)
	a.Logger.Info("health logic initiated")
}

func (a *App) initService() {
	a.Service = services.New(a.Config, a.Logger)
}
