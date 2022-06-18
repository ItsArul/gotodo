package cmd

import (
	"gotodo/cmd/controller"
	"gotodo/cmd/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run() {
  server := gin.New()
  server.Use(middleware.BasicAuth())

  router := server.Group("/v1/api")
  {
    router.GET("/", func(c *gin.Context) {
      c.JSON(http.StatusOK, "Welcome to TodoList")
    })
    router.GET("/todo", controller.FindAll)
    router.POST("/todo", controller.Create)
    router.PUT("/todo/:id", controller.Update)
    router.GET("/todo/:id", controller.FindById)
    router.DELETE("/todo/:id", controller.Delete)
  }

  server.Run(":3000")
}
