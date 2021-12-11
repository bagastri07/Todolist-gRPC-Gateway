package configs

import (
	"github.com/bagastri07/Todolist-gRPC-Gateway/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetDBConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./configs/todo.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Todo{})
	return db
}
