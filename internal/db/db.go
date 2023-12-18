package db

// TODO: Rip out Gorm. It's crap

import (
	"database/sql"
	"embed"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
)

var DB *sql.DB

//go:embed migrations/*.sql
var embedMigrations embed.FS

// const DB_FILE = ":memory:" // in memory great for development
const DB_FILE = "dev.db"

func init() {
	fmt.Println("Database initialising and migrating")
	db, err := sql.Open("sqlite3", DB_FILE)
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate
	DB = db

	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("sqlite3"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}

}
