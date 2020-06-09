package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"nowim.user/internal/config"
)

var db *gorm.DB

func init() {
	host := config.Config().Postgres.Host
	port := config.Config().Postgres.Port
	username := config.Config().Postgres.Username
	dbname := config.Config().Postgres.Database
	password := config.Config().Postgres.Password
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, username, dbname, password)

	var err error
	db, err = gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("database connect failed, err: %+v", err)
	}

	log.Info("database connected")
}

func DB() *gorm.DB {
	return db
}