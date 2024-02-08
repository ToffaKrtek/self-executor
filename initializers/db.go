package initializers

import (
	"os"

	"github.com/ToffaKrtek/self-executor/models"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open(os.Getenv("DB_DATABASE")), &gorm.Config{})
	if err != nil {
		panic("Ошибка подключения к Базе Данных")
	}
	migrate()
}

func migrate() {
	DB.AutoMigrate(
		models.User{},
		models.Script{},
		models.Permission{},
		models.Log{})
}
