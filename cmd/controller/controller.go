package controller

import (
	"context"
	"encoding/json"
	"gotodo/cmd/repository"
	"gotodo/cmd/entity"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func Create(c *gin.Context) {
  var todo entity.Todolist
  ctx := context.Background()

  t, err := repository.Create(ctx, todo)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H {
      "code": http.StatusBadRequest,
      "message": "cannot create todo",
      "error": err,
    })
  }

  c.Bind(&t)

  c.JSON(http.StatusOK, gin.H {
    "code": http.StatusOK,
    "message": "success create todo",
    "todo": t.Name,
  })
}

func Update(c *gin.Context) {
  var todo entity.Todolist
  ctx := context.Background()

  Id := c.Param("id")
  id, _ := strconv.Atoi(Id)
  update, _ := ioutil.ReadAll(c.Request.Body)
  json.Unmarshal(update, &todo)

  t, err := repository.Update(ctx, id, todo)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H {
      "code": http.StatusBadRequest,
      "message": "cannot update todo",
      "error": err,
    })
  }

  c.Bind(&todo)
  c.JSON(http.StatusOK, gin.H {
    "code": http.StatusOK,
    "message": "success update todo",
    "data": t,
  })
}

func FindById(c *gin.Context) {
  ctx := context.Background()
  Id := c.Param("id")
  id, _ := strconv.Atoi(Id)

  t, err := repository.FindById(ctx, id)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H {
      "code": http.StatusBadRequest,
      "message": "cannot find todo",
      "error": err,
    })
  }

  c.JSON(http.StatusOK, gin.H {
    "code": http.StatusOK,
    "message": "success update todo",
    "todoname": t.Name,
  })
}

func FindAll(c *gin.Context) {
  ctx := context.Background()

  t, err := repository.FindAll(ctx)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H {
      "code": http.StatusBadRequest,
      "message": "cannot find todo",
      "error": err,
    })
  }

  c.JSON(http.StatusOK, gin.H {
    "code": http.StatusOK,
    "message": "success update todo",
    "data": t,
  })
  
}
func Delete(c *gin.Context) {
  ctx := context.Background()

  Id := c.Param("id")
  id, _ := strconv.Atoi(Id)

  if err := repository.Delete(ctx, id) ; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "code": http.StatusBadRequest,
      "message": "cant delete data",
      "error": err,
    })
  }

  c.JSON(http.StatusOK, gin.H{
    "code": http.StatusOK,
    "message": "success delete data",
  })
}
