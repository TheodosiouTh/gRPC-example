package cmd

import (
	"context"
	"fmt"
	"strconv"
	"todo/client/messenger"
	"todo/todo"

	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Marks a task as done.",
	Long:  "Marks the task of the given id as done",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			fmt.Println("please enter a valid id.")
		}

		_, err = messenger.GetClient().Check(context.Background(), &todo.TaskId{Id: id})
		if err != nil {
			fmt.Printf("could not check task: %v", err)
		}
		fmt.Println("Task checked!")
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
