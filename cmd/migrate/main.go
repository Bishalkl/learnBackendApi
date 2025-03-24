package main

import (
	"log"
	"os"

	"github.com/bishalkl/learnBackendApi/config"
	"github.com/bishalkl/learnBackendApi/db"
	_ "github.com/go-sql-driver/mysql" // mysql driver
	mysqlDriver "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	mysqlMigrate "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// database connection
	cfg := mysqlDriver.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := db.NewMySQLStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// creating driver
	driver, err := mysqlMigrate.WithInstance(db, &mysqlMigrate.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// for migrating
	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Get current version and dirty state
	v, dirty, _ := m.Version()
	log.Printf("Version: %d, dirty: %v", v, dirty)

	// If the database is dirty, force a migration to fix it
	if dirty {
		log.Println("Database is in a dirty state. Running force to fix...")
		err := m.Force(int(v)) // Force the version to the last applied one
		if err != nil {
			log.Fatal("Error fixing dirty database version:", err)
		}
		log.Println("Database state cleaned successfully.")
	}

	// cli for up and down migrate
	cmd := os.Args[len(os.Args)-1]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}

}
