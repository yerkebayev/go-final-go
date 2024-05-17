package db

import (
	"fmt"
	"github.com/yerkebayev/go-final-go/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	username = "marat"
	password = "password"
	hostname = "127.0.0.1"
	port     = 5432
	database = "postgres"
	schema   = "test"
)

func Init() *gorm.DB {
	DSN := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?search_path=%s", username, password, hostname, port, database, schema)

	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.Teacher{}, &models.Student{}, &models.Attendance{}, &models.Course{}, &models.Session{}, &models.Image{})

	return db
}
