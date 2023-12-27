package session

import (
	"app/internal/db"

	"github.com/alexedwards/scs/sqlite3store"
	"github.com/alexedwards/scs/v2"
)

var Manager *scs.SessionManager

func InitSessionManager() {
	db := db.DB
	Manager = scs.New()
	Manager.Store = sqlite3store.New(db)
}
