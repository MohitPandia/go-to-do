package database

import (
	"database/sql"
	"fmt"
	"go-to-do/configs"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	PostgresDB *gorm.DB
	SQLDB      *sql.DB
)

func Connection() (*gorm.DB, *sql.DB) {
	fmt.Println("now came in sql_connection")
	var err error
	if PostgresDB != nil && SQLDB != nil {
		return PostgresDB, SQLDB
	}
	dsn := "host=" + configs.DB.Host +
		" user=" + configs.DB.Username +
		" password=" + configs.DB.Password +
		" dbname=" + configs.DB.Database +
		" port=" + configs.DB.Port +
		" sslmode=disable"

	PostgresDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("[Connection], Error in opening db")
	}

	SQLDB, err = PostgresDB.DB()
	if err != nil {
		log.Fatal("[Connection], Error in setting sqldb")
	}

	SQLDB.SetMaxIdleConns(10)
	SQLDB.SetMaxOpenConns(100)
	SQLDB.SetConnMaxLifetime(time.Hour)

	return PostgresDB, SQLDB
}
