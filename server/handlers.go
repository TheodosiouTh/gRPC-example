package main

import (
	"context"
	"todo/server/db"
	"todo/todo"
)

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
