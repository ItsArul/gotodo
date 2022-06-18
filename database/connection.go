package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type AppRun struct {

  Database struct {
    Host string
    Port string
    DBName string
    User string
    Password string
  }
}

var app *AppRun
var DB *gorm.DB

func GetConnection() *gorm.DB {
  config := initConfig()
  login := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4",
                      config.Database.User,
                      config.Database.Password,
                      config.Database.Host,
                      config.Database.Port,
                      config.Database.DBName )

  db, err := gorm.Open(mysql.Open(login), &gorm.Config{})
  if err != nil {
    fmt.Printf("can connect database")
    panic(err)
  }

  DB = db

  return db
}


func initConfig() *AppRun {
  if app == nil {
    app = getConfig()
  }

  return app
}

func getConfig() *AppRun {
  config := AppRun{}
  if err := godotenv.Load(); err != nil {
    config.Database.Host = "localhost"
    config.Database.Port = "3306"
    config.Database.User = "root"
    config.Database.Password = "12345"
    config.Database.DBName = "todolist"

    return &config
  }

  config.Database.Host = os.Getenv("DB_HOST")
  config.Database.Port = os.Getenv("DB_PORT")
  config.Database.User = os.Getenv("DB_USER")
  config.Database.Password = os.Getenv("DB_PASS")
  config.Database.DBName = os.Getenv("DB_NAME")

  return &config
}
