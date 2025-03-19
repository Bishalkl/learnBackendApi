package main

import (
	"database/sql"
	"log"

	"github.com/bishalkl/learnBackendApi/cmd/api"
	"github.com/bishalkl/learnBackendApi/config"
	"github.com/bishalkl/learnBackendApi/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	// connecting server with database
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	// check and connect the database
	initStorage(db)

	// creating server for apis
	server := api.NewAPIserver(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}

// func to check
func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB: Successfully connected")
}
