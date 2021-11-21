package cmd

import (
	"context"
	"fmt"
	"strings"
	"todo/client/messenger"
	"todo/todo"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to the task list",
	Run: func(cmd *cobra.Command, args []string) {
		taskName := strings.Join(args, " ")
		_, err := messenger.GetClient().Add(context.Background(), &todo.Task{Name: taskName})
		if err != nil {
			fmt.Printf("could not create task: %v", err)
		}
		fmt.Println("Task added!")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
