package cmd

import (
	"context"
	"fmt"
	"todo/client/messenger"
	"todo/todo"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := messenger.GetClient().List(context.Background(), &todo.Void{})
		if err != nil {
			fmt.Printf("could not find task: %v", err)
		}

		messenger.PrintTasks(tasks)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
