package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the application",
	Run: func(cmd *cobra.Command, args []string) {
		directories := viper.GetStringSlice("directories")
		if len(directories) == 0 {
			fmt.Println("No directories found. Please add directories using 'config add' command.")
			return
		}
		fmt.Println("Running application with directories:")
		for _, dir := range directories {
			fmt.Println("-", dir)
		}
		// Run the actual application logic here
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
