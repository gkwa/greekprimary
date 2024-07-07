package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration",
	Run: func(cmd *cobra.Command, args []string) {
		configFile := viper.ConfigFileUsed()
		if configFile != "" {
			fmt.Println(configFile)
		}
	},
}

var dirCmd = &cobra.Command{
	Use:   "dir",
	Short: "Manage configuration",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			err := cmd.Usage()
			if err != nil {
				fmt.Println("no directory path provided:", err)
			}
			return
		}

		configFile := viper.ConfigFileUsed()
		if configFile != "" {
			fmt.Println(configFile)
		}
	},
}

var addCmd = &cobra.Command{
	Use:   "add <directories>...",
	Short: "Add directories to the list",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide at least one directory to add")
			return
		}
		directories := viper.GetStringSlice("directories")
		for _, dir := range args {
			if !contains(directories, dir) {
				directories = append(directories, dir)
			}
		}
		viper.Set("directories", directories)
		if err := viper.WriteConfig(); err != nil {
			fmt.Println("Failed to update config file:", err)
		}
	},
}

var removeCmd = &cobra.Command{
	Use:   "remove <directories>...",
	Short: "Remove directories from the list",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide at least one directory to remove")
			return
		}
		directories := viper.GetStringSlice("directories")
		for _, dir := range args {
			if contains(directories, dir) {
				directories = remove(directories, dir)
			} else {
				fmt.Println("Directory not found:", dir)
				fmt.Println("Current directories:")
				showDirs()
			}
		}
		viper.Set("directories", directories)
		if err := viper.WriteConfig(); err != nil {
			fmt.Println("Failed to update config file:", err)
		}
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all stored directories",
	Run: func(cmd *cobra.Command, args []string) {
		showDirs()
	},
}

func showDirs() {
	directories := viper.GetStringSlice("directories")
	if len(directories) == 0 {
		fmt.Println("No directories found")
	} else {
		for _, dir := range directories {
			fmt.Println("-", dir)
		}
	}
}

func init() {
	dirCmd.AddCommand(addCmd)
	dirCmd.AddCommand(removeCmd)
	dirCmd.AddCommand(listCmd)
	configCmd.AddCommand(dirCmd)
	rootCmd.AddCommand(configCmd)
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if strings.EqualFold(s, item) {
			return true
		}
	}
	return false
}

func remove(slice []string, item string) []string {
	for i, s := range slice {
		if strings.EqualFold(s, item) {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
