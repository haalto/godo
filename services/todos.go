package services

import (
	"github.com/haalto/godo/data"
	"github.com/haalto/godo/models"
)

func GetTodos() models.Todos {
	var resultArray models.Todos
	db := data.GetDB()
	var todo models.Todo
	err := db.Model(todo).Find(&resultArray).Error

	if err != nil {
		return nil
	}

	return resultArray
}
