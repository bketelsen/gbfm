package admin

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
	"github.com/qor/auth"
)

// auth0Provider is a Provider implementation that uses Auth0
type auth0Provider struct {
	provider goth.Provider
}

func newAuth0Provider(ghKey, ghSecret, host string) auth.Provider {
	ghProvider := github.New(ghKey, ghSecret, fmt.Sprintf("%s/auth/github/callback", host))
	goth.UseProviders(ghProvider)
	return &auth0Provider{
		provider: github.New(ghKey, ghSecret, fmt.Sprintf("%s/auth/github/callback", host)),
	}
}

func (ap auth0Provider) GetName() string {
	return "auth0"
}

func (ap auth0Provider) ConfigAuth(a *auth.Auth) {
	// TODO I think
}

func (ap auth0Provider) Login(c *auth.Context) {
	// TODO I think
}

func (ap auth0Provider) Logout(c *auth.Context) {
	// TODO I think
}

func (ap auth0Provider) Register(c *auth.Context) {
	// TODO I think
}

func (ap auth0Provider) Callback(c *auth.Context) {
	// TODO I think
}

func (ap auth0Provider) ServeHTTP(c *auth.Context) {
	router := mux.NewRouter()
	router.HandleFunc("/auth/{provider}/callback", func(w http.ResponseWriter, r *http.Request) {
		user, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			fmt.Fprintln(w, r)
			return
		}
		// t, _ := template.New("foo").Parse(userTemplate)
		// t.Execute(res, user)
	}).Methods("GET")

	router.HandleFunc("/logout/{provider}", func(w http.ResponseWriter, r *http.Request) {
		gothic.Logout(w, r)
		w.Header().Set("Location", "/")
		w.WriteHeader(http.StatusTemporaryRedirect)
	}).Methods("GET")

	router.HandleFunc("/auth/{provider}", func(w http.ResponseWriter, r *http.Request) {
		if gothUser, err := gothic.CompleteUserAuth(w, r); err == nil {
			// t, _ := template.New("foo").Parse(userTemplate)
			// t.Execute(res, gothUser)
		} else {
			gothic.BeginAuthHandler(w, r)
		}

	})

	router.ServeHTTP(c.Writer, c.Request)
}
