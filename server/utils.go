package main

import (
	"todo/server/db"
	"todo/todo"
)

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
