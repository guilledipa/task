package cmd

import (
	"fmt"
	"log"
	"os"
	"proyectos/task/db"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.Flags().StringVarP(&dbPath, "db_path", "d", "~/tasks.db", "Database path")
	cobra.OnInitialize(initDB)
}

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "task is a CLI for managing your TODOs.",
	Long: `A gophercises excercise #7:
	https://github.com/gophercises/task`,
}

var dbPath string

// Execute will run rootCmd after being called
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Root cmd: %v.", err)
	}
}

func initDB() {
	if err := db.Init(dbPath); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
