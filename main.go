package main

import (
	"app/internal/todo"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func main() {

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	DB = db

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./web/public"))

	// Note to self: Routing in Go Stdlib is awkward :(
	mux.HandleFunc("/", todo.Handler)

	mux.Handle("/public/", http.StripPrefix("/public/", fs))
	log.Print("Listening on http://localhost:3000")
	err = http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal(err)
	}

}
