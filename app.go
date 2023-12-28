package main

import (
	"app/domain/todo"
	"app/domain/user"
	_ "app/internal/db"
	"app/internal/session"
	"log"
	"net/http"
	"strings"

	"app/web/pages"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	fs := http.FileServer(http.Dir("./web/public")) // TODO: Embed public
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))
	session.InitSessionManager()
	// router.Use(user.AuthMiddleWare([]string{
	// 	"/user/login",
	// 	"/public",
	// }))
	router.Use(user.RequireAuthForPathMiddleware(func(urlPath string) bool {
		if strings.HasPrefix(urlPath, "/todo") {
			return true
		}
		return false
	}))

	todo.RegisterRoutes(router.PathPrefix("/todo").Subrouter())
	user.RegisterRoutes(router.PathPrefix("/user").Subrouter())
	router.Handle("/", templ.Handler(pages.IndexPage()))

	log.Print("Listening on http://localhost:3000")
	//err := http.ListenAndServe(":3000", router)
	err := http.ListenAndServe(":3000", session.Manager.LoadAndSave(router))
	if err != nil {
		log.Fatal(err)
	}

}
