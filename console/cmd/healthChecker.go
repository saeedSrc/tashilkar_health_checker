package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"log"
	"tashilkar_health_checker/config"
	"tashilkar_health_checker/logic"
	"tashilkar_health_checker/repo"
	services "tashilkar_health_checker/service"
	"time"
)

var HealthChecker = &cobra.Command{
	Use:   "health_checker",
	Short: "health checker command",
	Long:  `a background job for checking registered apis withing their interval`,
	Run: func(cmd *cobra.Command, args []string) {
		logger, err := zap.NewProduction()
		if err != nil {
			log.Fatal(err)
		}
		sugar := logger.Sugar()
		sugar.Info("health checker command has started...")
		cfg := config.Init("./config.yaml")
		mongo := repo.NewDB(sugar, cfg)
		mongoConn := mongo.Init()
		healthRepo := repo.NewHealthCheckerRepo(mongoConn, sugar, cfg)
		service := services.New(cfg, sugar)
		healthLogic := logic.NewHealthCheckerLogic(healthRepo, sugar, service, cfg)
		for true {
			err = healthLogic.Check()
			if err != nil {
				sugar.Errorf("there is an error in checking api heath: %v", err)
			}
			time.Sleep(10 * time.Second)
		}

	},
}

func init() {
	rootCmd.AddCommand(HealthChecker)
}
