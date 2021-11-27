package messenger

import (
	"fmt"
	"log"
	"todo/todo"

	"google.golang.org/grpc"
)

func GetClient() todo.TodoClient {

	connection, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to BE: %v", err)
	}

	return todo.NewTodoClient(connection)
}

func PrintTasks(list *todo.Tasks) {
	for _, task := range list.Tasks {
		PrintTask(task)
	}
}

func PrintTask(task *todo.Task) {
	var prefix string
	if task.Done {
		prefix = "[X]"
	} else {
		prefix = "[ ]"
	}

	fmt.Printf("%s %s (id: %d)\n", prefix, task.Name, task.Id)
}
