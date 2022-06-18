package repository

import (
	"context"
	"fmt"
	"gotodo/cmd/entity"
	"gotodo/database"
)

func Create(ctx context.Context, todo entity.Todolist) (entity.Todolist, error) {
  if err := database.DB.WithContext(ctx).Create(&todo).Error; err != nil {
    fmt.Printf("cannot create todo")
    panic(err)
  }

  return todo, nil
  
}

func Update(ctx context.Context, id int, todo entity.Todolist) (entity.Todolist, error) {
  var t entity.Todolist
  if err := database.DB.WithContext(ctx).Model(&t).Where("id = ?", id).Updates(&todo).Error; err != nil {
    fmt.Printf("cannot update")
    panic(err)
  }

  return todo, nil
}

func FindById(ctx context.Context, id int) (entity.Todolist, error) {
  var todo entity.Todolist
  if err := database.DB.WithContext(ctx).First(&todo, id).Error; err != nil {
    fmt.Printf("cannot find")
    panic(err)
  }

  return todo, nil
}

func FindAll(ctx context.Context) ([]entity.Todolist, error) {
  var todo []entity.Todolist
  if err := database.DB.WithContext(ctx).Find(&todo).Error; err != nil {
    fmt.Printf("cannot find all")
    panic(err)
  }

  return todo, nil
}


func Delete(ctx context.Context, id int) error {
  var todo entity.Todolist

  if err := database.DB.WithContext(ctx).Where("id = ?", id).Delete(todo).Error; err != nil {
    fmt.Printf("cannot delete")
    panic(err)
  }

  return nil
}
