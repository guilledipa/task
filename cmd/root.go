package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/guilledipa/task/db"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.Flags().StringVarP(&dbPath, "db_path", "d", "tasks.db", "Database path")
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
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dbAtHome := filepath.Join(home, dbPath)
	if err := db.Init(dbAtHome); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
