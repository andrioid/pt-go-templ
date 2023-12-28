package user

import (
	"app/internal/db"
	"app/internal/session"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type UserModel struct {
	ID         int64
	Email      string
	IsAdmin    bool
	IsVerified bool
	// Private. There should be no reason for reading outside user
	passwd   string
	hashType string
}

func GetUser(ctx context.Context, email string) (UserModel, error) {
	user := UserModel{}
	res := db.DB.QueryRowContext(ctx, "SELECT id, email, passwd, hash_type, is_email_verified, is_admin FROM user WHERE email=?", email)
	err := res.Scan(&user.ID, &user.Email, &user.passwd, &user.hashType, &user.IsVerified, &user.IsAdmin)
	if err != nil {
		return user, err
	}
	return user, nil
}

func LoginFormHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("wat")
	// TODO: Check referrer
	// TODO: Check user
	user := r.FormValue("user")
	password := r.FormValue("password")
	err := Login(r.Context(), user, password)
	if err == nil {
		// Login succeded, redirecting
		w.Header().Add("HX-Redirect", "/") // HTMX specific
		return
	}

	component := LoginForm(err)
	component.Render(r.Context(), w)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if err := Logout(r.Context()); err != nil {
		log.Println(err)
		http.Error(w, "Failed to log out", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func RequireAuthForPathMiddleware(requireAuthFunc func(string) bool) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO: Check session

			if requireAuthFunc(r.URL.Path) && !IsLoggedIn(r.Context()) {
				referer := r.URL.Path
				http.Redirect(w, r, fmt.Sprintf("/user/login/?ref=%s", referer), http.StatusFound)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func IsLoggedIn(ctx context.Context) bool {
	if session.Manager == nil {
		return false
	}
	userId := session.Manager.GetString(ctx, "user_id")
	return userId != "" // False if empty
}
