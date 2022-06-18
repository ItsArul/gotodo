package entity

import "time"

type Todolist struct {
  ID int `json:"id" form:"id" binding:"required"`
  Name string `json:"name" form:"name" binding:"required"`
  Time time.Time `json:"time" form:"time"`
}
