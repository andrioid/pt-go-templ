package main

import (
	_ "app/internal/db"
	"app/internal/todo"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./web/public"))

	// Note to self: Routing in Go Stdlib is awkward :(
	mux.HandleFunc("/", todo.Handler)

	mux.Handle("/public/", http.StripPrefix("/public/", fs))
	log.Print("Listening on http://localhost:3000")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal(err)
	}

}
