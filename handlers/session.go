package handlers

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func MyHandler(w http.ResponseWriter, r *http.Request) {
	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.
	session, _ := store.Get(r, "mbdg_session")
	// Set some session values.
	session.Values["foo"] = "bar"
	session.Values[42] = 43
	// Save it.
	session.Save(r, w)
}
