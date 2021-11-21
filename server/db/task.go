package db

import (
	"gRPC-example/todo"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Id   uint
	Name string
	Done bool
}

func FindAll() []Task {
	var tasks []Task
	db.Find(&tasks)
	return tasks
}

func FindById(taskId *todo.TaskId) Task {
	var task Task
	db.Find(&task, taskId)
	return task
}

func Create(input *todo.Task) Task {
	taskToCreate := Task{
		Name: input.Name,
		Done: false,
	}

	db.Create(&taskToCreate)

	return FindById(&todo.TaskId{Id: uint64(taskToCreate.Id)})
}

func Check(taskId *todo.TaskId) Task {
	task := FindById(taskId)

	task.Done = true
	db.Save(&task)

	return FindById(&todo.TaskId{Id: uint64(task.Id)})
}

func Remove(task *todo.TaskId) []Task {
	db.Delete(&Task{}, task.Id)
	return FindAll()
}
