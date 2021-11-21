package cmd

import (
	"context"
	"fmt"
	"strconv"
	"todo/client/messenger"
	"todo/todo"

	"github.com/spf13/cobra"
)

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Finds a specific taks",
	Long:  "Finds a specific task based on a given id.",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			fmt.Println("please enter a valid id.")
		}

		task, err := messenger.GetClient().Find(context.Background(), &todo.TaskId{Id: id})
		if err != nil {
			fmt.Printf("could not find task: %v", err)
		}

		messenger.PrintTask(task)
	},
}

func init() {
	rootCmd.AddCommand(findCmd)
}
