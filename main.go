package main

import (
	"gotodo/cmd"
	"gotodo/cmd/entity"
	"gotodo/database"
)

func main() {
	db := database.GetConnection()
  	db.AutoMigrate(&entity.Todolist{})
  	cmd.Run()
}

