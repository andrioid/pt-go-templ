package main

import (
	_ "app/internal/db"
	"app/internal/todo"
	"app/internal/user"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	fs := http.FileServer(http.Dir("./web/public"))

	todo.RegisterRoutes(router.PathPrefix("/todo").Subrouter())
	user.RegisterRoutes(router.PathPrefix("/user").Subrouter())

	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))
	log.Print("Listening on http://localhost:3000")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal(err)
	}

}
