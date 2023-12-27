package user

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// TODO: Login page

func RootHandler(w http.ResponseWriter, r *http.Request) {
	component := LoginPage(nil)
	component.Render(r.Context(), w)
}

func LoginFormHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Check referrer
	// TODO: Check user
	r.ParseForm()
	user := r.FormValue("user")
	password := r.FormValue("password")
	fmt.Printf("%v\n", r.Form)
	log.Printf("user: '%s', pass: '%s'", user, password)
	w.Header().Add("Content-Type", "text/html")
	var err error
	if user == "admin" && password == "password" {
		log.Println("Login good")
		// "secure" defaults, lol
		w.Header().Add("HX-Redirect", "/") // Tell HTMX to redirect

		//http.Redirect(w, r, "/user/", http.StatusFound)
	} else {
		log.Println("Login bad", user, password)
		err = errors.New("invalid login")
	}

	w.WriteHeader(http.StatusOK)
	component := LoginForm(err)
	component.Render(r.Context(), w)
}

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/", RootHandler).Methods("GET")
	r.HandleFunc("/", LoginFormHandler).Methods("POST")
}
