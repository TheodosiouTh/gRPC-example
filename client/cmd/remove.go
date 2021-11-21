package cmd

import (
	"context"
	"fmt"
	"strconv"
	"todo/client/messenger"
	"todo/todo"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes a task.",
	Long:  "Removes the task with the giver id.",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			fmt.Println("please enter a valid id.")
		}

		_, err = messenger.GetClient().Delete(context.Background(), &todo.TaskId{Id: id})
		if err != nil {
			fmt.Printf("could not delete task: %v", err)
		}
		fmt.Println("Task deleted!")
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
