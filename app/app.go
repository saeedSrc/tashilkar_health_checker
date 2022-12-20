package app

import (
	"go.uber.org/zap"
	"log"
	"tashilkar_health_checker/repo"
)

type App struct {
	Logger *zap.SugaredLogger
}

func NewApp() *App {
	app := &App{}
	return app
}

func (a *App) Init() {
	a.initialLogger()
	a.initRepo(a.Logger)
}

func (a *App) initialLogger() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	a.Logger = logger.Sugar()
}

func (a *App) initRepo(l *zap.SugaredLogger) {
	repo.Init(l)
}
