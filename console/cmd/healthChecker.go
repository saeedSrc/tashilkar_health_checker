package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"tashilkar_health_checker/logic"
)

var DailyDealsUserUpdater = &cobra.Command{
	Use:   "health_checker",
	Short: "health checker command",
	Long:  `a background job for checking registered apis withing their interval`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("running health checker ....")
		healthLogic := logic.NewHealthCheckerLogic()
		err := healthLogic.Check()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(DailyDealsUserUpdater)
}
