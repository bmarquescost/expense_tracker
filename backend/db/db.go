package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	
	// username:password@protocol(address)/dbname?param=value
	dsn := fmt.Sprintf("%s:%s@tcp(mysql-db:3306)/%s?parseTime=true", dbUser, dbPassword, dbName)
	log.Println(dsn)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("DB is not reachable:", err)
		return nil, err
	}

	log.Println("Connected to MySQL database")
	return db, nil
}