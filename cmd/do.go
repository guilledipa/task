package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/guilledipa/task/db"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(doCmd)
}

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task (passing id number) as complete.",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("do: failed to parse %v argument, %v", arg, err)
				continue
			}
			ids = append(ids, id)
		}
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("do something went wrong:", err)
			os.Exit(1)
		}
		for _, id := range ids {
			if id < 1 || id > len(tasks) {
				fmt.Println("Invalid task number:", id)
				continue
			}
			task := tasks[id-1]
			if err := db.DeleteTask(task.Key); err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed. Error: %s\n", id, err)
				continue
			}
			fmt.Printf("Marked \"%d\" as completed.\n", id)
		}
	},
}
