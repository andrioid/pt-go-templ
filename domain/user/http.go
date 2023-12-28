package user

import (
	"github.com/a-h/templ"
	"github.com/gorilla/mux"
)

// /user prefix
func RegisterRoutes(r *mux.Router) {
	r.Handle("/login/", templ.Handler(LoginPage())).Methods("GET")
	r.HandleFunc("/login/", LoginFormHandler).Methods("POST")
	r.HandleFunc("/logout/", LogoutHandler).Methods("GET")
}
