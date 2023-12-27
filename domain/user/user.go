package user

import (
	"app/internal/session"
	"errors"
	"log"
	"net/http"

	"github.com/a-h/templ"
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
	user := r.FormValue("user")
	password := r.FormValue("password")
	var err error
	if user == "admin" && password == "password" {
		session.Manager.Put(r.Context(), "user_id", "123")
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
	r.Handle("/", templ.Handler(LoginPage(nil))).Methods("GET")
	//r.HandleFunc("/", RootHandler).Methods("GET")
	r.HandleFunc("/", LoginFormHandler).Methods("POST")
}
