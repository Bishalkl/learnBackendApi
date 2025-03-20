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
	dbConnection, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	defer dbConnection.Close() //Make sure the DB connection is closed

	// Check the DB connection
	if err := initStorage(dbConnection); err != nil {
		log.Fatal("Error initializing database: ", err)
	}

	// create an start the server
	server := api.NewAPIserver(":8080", dbConnection)
	if err := server.Run(); err != nil {
		log.Fatal("Error running server: ", err)
	}

}

// initStorage checks the database connection and ensure it's reachable
func initStorage(db *sql.DB) error {
	if err := db.Ping(); err != nil {
		return err
	}
	log.Println("DB: Successfully connected")
	return nil
}
