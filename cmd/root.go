package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "task is a CLI for managing your TODOs.",
	Long: `A gophercises excercise #7:
	https://github.com/gophercises/task`,
}

var dbPath string

// Execute will run rootCmd after being called
func Execute() {
	rootCmd.Flags().StringVarP(&dbPath, "db_path", "s", "~/tasks.db", "Database path")
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Root cmd: %v.", err)
	}
}
