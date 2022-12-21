package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"log"
	"tashilkar_health_checker/config"
	"tashilkar_health_checker/logic"
	"tashilkar_health_checker/repo"
	services "tashilkar_health_checker/service"
)

var DailyDealsUserUpdater = &cobra.Command{
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
		mongo := repo.Init(sugar)
		cfg := config.Init("./config.yaml")
		healthRepo := repo.NewHealthCheckerRepo(mongo, sugar)
		healthLogic := logic.NewHealthCheckerLogic(healthRepo, sugar, services.New(cfg, sugar), cfg)
		err = healthLogic.Check()
		if err != nil {
			sugar.Errorf("there is an error in checking api' heath: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(DailyDealsUserUpdater)
}
