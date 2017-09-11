package models

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))

// GetSession to return the session
func GetSession(r *http.Request) *sessions.Session {
	session, err := store.Get(r, "session")
	if err != nil {
		panic(err)
	}
	return session
}

// AllSessions function to return all the sessions
func AllSessions(r *http.Request) (interface{}, interface{}) {
	session := GetSession(r)
	id := session.Values["id"]
	username := session.Values["username"]

	return id, username
}
