package cmd

import (
	"github.com/spf13/cobra"
)

type Command struct {
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "datacmd",
	Short: "starting point of all other commands.",
	Long:  `starting point of all other commands.`,
}

//Execute execute cmd
func init() {
}

//Execute execute
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		return
	}
}
