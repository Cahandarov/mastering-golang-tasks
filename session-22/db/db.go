package db

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"session-22/model"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

var DB *gorm.DB

func InitDB() error {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Baku",
		host, port, user, password, dbname,
	)

	log.Info("Starting database connection: ", host, port)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatal("ActionLog.DbInit.error ", err)
		return err
	}
	DB = db
	err = Migrate()

	if err != nil {
		return err
	}
	return nil
}
func GetDB() *gorm.DB {
	return DB
}
func Migrate() error {
	log.Info("Migration database started: ", host, port)
	err := GetDB().AutoMigrate(&model.Book{})

	if err != nil {
		log.Fatal("ActionLog.Migrate.error ", err.Error())
		return err
	}
	return nil
}
