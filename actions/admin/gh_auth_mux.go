package admin

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/markbates/goth/gothic"
)

func addGHAuthRoutes(mx *http.ServeMux) {
	router := mux.NewRouter()

	// login start
	router.HandleFunc(loginPath, func(w http.ResponseWriter, r *http.Request) {
		u, err := gothic.GetAuthURL(w, r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err)
			return
		}
		http.Redirect(w, r, u, http.StatusTemporaryRedirect)
	})

	// GH callback
	router.HandleFunc(callbackPath, func(w http.ResponseWriter, r *http.Request) {
		gothUsr, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			fmt.Fprintln(w, r)
			return
		}
		if err := gothic.StoreInSession(userIDKey, gothUsr.Name, r, w); err != nil {
			fmt.Fprintf(w, "error: %s", err)
			return
		}
	}).Methods("GET")

	// logout
	router.HandleFunc(logoutPath, func(w http.ResponseWriter, r *http.Request) {
		gothic.Logout(w, r)
		w.Header().Set("Location", "/admin")
		w.WriteHeader(http.StatusTemporaryRedirect)
	}).Methods("GET")

	mx.Handle("/", router)
}
