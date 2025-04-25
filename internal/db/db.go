package db

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var DB *gorm.DB

func InitDB() {
	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "postgres"
	dbPass := "Asilzhan7"
	dbName := "postgres"
	sslmode := "disable"

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPass, dbHost, dbPort, dbName, sslmode)

	sqlDB, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("[DB ERROR] SQL Open failed:", err)
	}

	driver, err := migratepg.WithInstance(sqlDB, &migratepg.Config{})
	if err != nil {
		log.Fatal("[DB ERROR] Migrate driver init failed:", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/db/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal("[DB ERROR] Migrate init failed:", err)
	}

	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatal("[DB ERROR] Migration failed:", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("[DB ERROR] GORM Open failed:", err)
	}

	DB = gormDB

}
