package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func dbConfig() string {
	db_name := os.Getenv("DB_NAME")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", db_user, db_password, db_host, db_port, db_name)
}

func DbConnect() *sql.DB {

	dbpath := dbConfig()
	db_connect, err := sql.Open("mysql", dbpath)
	if err != nil {
		log.Println("error al conectar a la base de datos")
		log.Println("Error", err)
	}
	return db_connect
}

func DbTest() {
	dbconect := DbConnect()
	err := dbconect.Ping()
	if err != nil {
		log.Println("Error: ", err)
	} else {
		fmt.Println("Pong Exitoso")
	}
}
