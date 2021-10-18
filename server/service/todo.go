package service

import (
	"server/database"
	"server/entity"
)

func FindAllTodos() []entity.Todo {
	todos := []entity.Todo{}
	db := database.Connect()
	db.Order("ID asc").Find(&todos)
	defer db.Close()

	return todos
}
