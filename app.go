package main

import (
	"app/domain/todo"
	"app/domain/user"
	_ "app/internal/db"
	"app/internal/session"
	"log"
	"net/http"

	"app/web/pages"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	fs := http.FileServer(http.Dir("./web/public"))
	session.InitSessionManager()

	todo.RegisterRoutes(router.PathPrefix("/todo").Subrouter())
	user.RegisterRoutes(router.PathPrefix("/user").Subrouter())
	router.Handle("/", templ.Handler(pages.IndexPage()))

	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))
	log.Print("Listening on http://localhost:3000")
	//err := http.ListenAndServe(":3000", router)
	err := http.ListenAndServe(":3000", session.Manager.LoadAndSave(router))
	if err != nil {
		log.Fatal(err)
	}

}
