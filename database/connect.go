package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

func Connect(logger *log.Logger) (*sql.DB, error) {
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlDb := os.Getenv("MYSQL_DB")

	if mysqlHost == "" {
		return nil, fmt.Errorf("No MYSQL_HOST provided")
	}
	if mysqlUser == "" {
		return nil, fmt.Errorf("No MYSQL_USER provided")
	}
	if mysqlPassword == "" {
		return nil, fmt.Errorf("No MYSQL_PASSWORD provided")
	}
	if mysqlDb == "" {
		return nil, fmt.Errorf("No MYSQL_DB provided")
	}

	connStr := fmt.Sprintf("%s:%s@/%s", mysqlUser, mysqlPassword, mysqlDb)

	db, err := sql.Open(mysqlHost, connStr)

	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
