package main

import (
	"context"
	"fmt"
	"gRPC-example/server/db"
	"gRPC-example/todo"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	err := db.Init()
	if err != nil {
		log.Fatalf("Could not connectto the Database: %v", err)
	}

	errorMessage := initializeServer()
	if errorMessage != "" {
		log.Fatalf(errorMessage)
	}

}

type todoServer struct {
	todo.UnimplementedTodoServer
}

const PORT = ":8888"

func initializeServer() string {

	server := grpc.NewServer()

	var todoService todoServer
	todo.RegisterTodoServer(server, todoService)

	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		return fmt.Sprintf("could not listen to port %s: %v", PORT, err)
	}
	log.Printf("listening to port %s", PORT)

	err = server.Serve(listener)
	if err != nil {
		return fmt.Sprintf("Could not serve: %v", err)
	}
	return ""
}

func (server todoServer) List(context.Context, *todo.Void) (*todo.Tasks, error) {
	foundTasks := db.FindAll()
	tasksToReturn := convertTasksToMessage(foundTasks)
	return tasksToReturn, nil
}

func (server todoServer) Find(ctx context.Context, todoId *todo.TaskId) (*todo.Task, error) {
	foundItem := db.FindById(todoId)
	return convertTaskToMessage(foundItem), nil
}

func (server todoServer) Add(ctx context.Context, task *todo.Task) (*todo.Task, error) {
	createdTask := db.Create(task)
	return convertTaskToMessage(createdTask), nil
}

func (server todoServer) Check(ctx context.Context, taskId *todo.TaskId) (*todo.Task, error) {
	updatedTask := db.Check(taskId)
	itemsToReturn := convertTaskToMessage(updatedTask)
	return itemsToReturn, nil
}

func (server todoServer) Delete(ctx context.Context, taskId *todo.TaskId) (*todo.Tasks, error) {
	foundTasks := db.Remove(taskId)
	itemsToReturn := convertTasksToMessage(foundTasks)
	return itemsToReturn, nil
}

/* HELPER FUNCTIONS*/
func convertTasksToMessage(tasksToTransform []db.Task) *todo.Tasks {
	var list todo.Tasks
	for _, task := range tasksToTransform {
		list.Tasks = append(list.Tasks, convertTaskToMessage(task))
	}
	return &list
}

func convertTaskToMessage(taskToTransform db.Task) *todo.Task {
	return &todo.Task{
		Id:   uint64(taskToTransform.Id),
		Name: taskToTransform.Name,
		Done: taskToTransform.Done,
	}
}
