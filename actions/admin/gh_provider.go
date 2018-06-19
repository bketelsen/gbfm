package admin

import (
	"fmt"
	"net/http"
	// "time"

	"github.com/gorilla/mux"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
	"github.com/qor/admin"
	"github.com/qor/auth"
	"github.com/qor/auth/claims"
	"github.com/qor/qor"
)

const (
	loginPath    = "/admin/auth/login?provider=github"
	logoutPath   = "/admin/auth/logout?provider=github"
	callbackPath = "/admin/auth/callback?provider=github"
	userIDKey    = "user_id"
)

type providerAuther interface {
	auth.Provider
	admin.Auth
}

// ghProvider is a Provider implementation that uses Auth0
// it also implements admin.Auth.
//
// don't create this directly, use newGHProvider instead. it needs
// to set up global state
type ghProvider struct{}

func newGHProvider(ghKey, ghSecret, host string) providerAuther {
	prov := github.New(ghKey, ghSecret, fmt.Sprintf("%s%s", host, callbackPath))
	goth.UseProviders(prov)
	return &ghProvider{}
}

func (ap ghProvider) GetName() string {
	return "auth0"
}

func (ap ghProvider) ConfigAuth(a *auth.Auth) {
	a.Config.LoginHandler = func(c *auth.Context, fn func(*auth.Context) (*claims.Claims, error)) {

	}
	a.Config.LogoutHandler = ap.Logout
}

// /login/github
func (ap ghProvider) Login(c *auth.Context) {
}

// /logout/github
func (ap ghProvider) Logout(c *auth.Context) {
}

func (ap ghProvider) Register(c *auth.Context) {
}

func (ap ghProvider) Callback(c *auth.Context) {
}

func (ap ghProvider) ServeHTTP(c *auth.Context) {
	router := mux.NewRouter()

	// login start
	router.HandleFunc(loginPath, func(w http.ResponseWriter, r *http.Request) {
		u, err := gothic.GetAuthURL(c.Writer, c.Request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(c.Writer, err)
			return
		}
		http.Redirect(c.Writer, c.Request, u, http.StatusTemporaryRedirect)
	})

	// GH callback
	router.HandleFunc(callbackPath, func(w http.ResponseWriter, r *http.Request) {
		_, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			fmt.Fprintln(w, r)
			return
		}
	}).Methods("GET")

	// logout
	router.HandleFunc(logoutPath, func(w http.ResponseWriter, r *http.Request) {
		gothic.Logout(w, r)
		w.Header().Set("Location", "/admin")
		w.WriteHeader(http.StatusTemporaryRedirect)
	}).Methods("GET")

	router.ServeHTTP(c.Writer, c.Request)
}

func (ap ghProvider) LoginURL(c *admin.Context) string {
	fmt.Println("login url!")
	return "/admin/auth/login?provider=github"
}

func (ap ghProvider) LogoutURL(c *admin.Context) string {
	fmt.Println("logout url!")
	return "/admin/auth/logout?provider=github"
}

func (ap ghProvider) GetCurrentUser(c *admin.Context) qor.CurrentUser {
	fmt.Println("get current user!")
	gothUser, err := gothic.GetFromSession(userIDKey, c.Request)
	if err != nil {
		authURL, err := gothic.GetAuthURL(c.Writer, c.Request)
		if err != nil {
			return nil
		}
		http.Redirect(c.Writer, c.Request, authURL, http.StatusTemporaryRedirect)
		return nil
	}
	return qorCurrentUser{displayName: gothUser}
}
