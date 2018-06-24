package admin

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
	"github.com/qor/admin"
	"github.com/qor/qor"
)

func init() {
	// always using github as the auth provider
	gothic.GetProviderName = func(*http.Request) (string, error) {
		return "github", nil
	}
}

const (
	loginPath    = "/auth/login"
	logoutPath   = "/auth/logout"
	callbackPath = "/auth/callback"
	userIDKey    = "user_id"
)

// type providerAuther interface {
// 	auth.Provider
// 	admin.Auth
// }

// ghProvider is an admin.Auth implementation that uses GH
//
// don't create this directly, use newGHProvider instead. it needs
// to set up global state
type ghProvider struct{}

func newGHProvider(ghKey, ghSecret, host string) *ghProvider {
	callbackURL := "http://localhost:9000" + callbackPath
	prov := github.New(ghKey, ghSecret, callbackURL)
	goth.UseProviders(prov)
	return &ghProvider{}
}

// func (ap ghProvider) GetName() string {
// 	return "auth0"
// }

// func (ap ghProvider) ConfigAuth(a *auth.Auth) {
// 	a.Config.LoginHandler = func(c *auth.Context, fn func(*auth.Context) (*claims.Claims, error)) {

// 	}
// 	a.Config.LogoutHandler = ap.Logout
// }

// // /login/github
// func (ap ghProvider) Login(c *auth.Context) {
// }

// // /logout/github
// func (ap ghProvider) Logout(c *auth.Context) {
// }

// func (ap ghProvider) Register(c *auth.Context) {
// }

// func (ap ghProvider) Callback(c *auth.Context) {
// }

// func (ap *ghProvider) ServeHTTP(c *auth.Context) {
// 	router := mux.NewRouter()

// 	// login start
// 	router.HandleFunc(loginPath, func(w http.ResponseWriter, r *http.Request) {
// 		u, err := gothic.GetAuthURL(c.Writer, c.Request)
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			fmt.Fprintln(c.Writer, err)
// 			return
// 		}
// 		http.Redirect(c.Writer, c.Request, u, http.StatusTemporaryRedirect)
// 	})

// 	// GH callback
// 	router.HandleFunc(callbackPath, func(w http.ResponseWriter, r *http.Request) {
// 		gothUsr, err := gothic.CompleteUserAuth(w, r)
// 		if err != nil {
// 			fmt.Fprintln(w, r)
// 			return
// 		}
// 		if err := gothic.StoreInSession(userIDKey, gothUsr.Name, r, w); err != nil {
// 			fmt.Fprintf(w, "error: %s", err)
// 			return
// 		}
// 	}).Methods("GET")

// 	// logout
// 	router.HandleFunc(logoutPath, func(w http.ResponseWriter, r *http.Request) {
// 		gothic.Logout(w, r)
// 		w.Header().Set("Location", "/admin")
// 		w.WriteHeader(http.StatusTemporaryRedirect)
// 	}).Methods("GET")

// 	// TODO: handle failure callback paths here

// 	router.ServeHTTP(c.Writer, c.Request)
// }

func (ap ghProvider) LoginURL(c *admin.Context) string {
	loginPath, err := gothic.GetAuthURL(c.Writer, c.Request)
	if err != nil {
		return ""
	}
	fmt.Println("LoginURL returning", loginPath)
	return loginPath
}

func (ap ghProvider) LogoutURL(c *admin.Context) string {
	return logoutPath
}

func (ap ghProvider) GetCurrentUser(c *admin.Context) qor.CurrentUser {
	path := c.Request.URL.Path
	fmt.Println("get current user with path", path)
	if strings.HasPrefix(path, loginPath) || strings.HasPrefix(path, callbackPath) {
		fmt.Println("get current user exception")
		return qorCurrentUser{displayName: "_inprogress"}
	}
	gothUser, err := gothic.GetFromSession(userIDKey, c.Request)
	if gothUser == "" {
		fmt.Println("no user in goth session, trying to redirect to auth", err)
		if _, err := gothic.GetAuthURL(c.Writer, c.Request); err != nil {
			if err != nil {
				fmt.Println("get current user error getting goth auth url", err)
				return nil
			}
			fmt.Println("get current user error getting user from session", err)
			return nil
		}
	}
	if err != nil {
		fmt.Println("error getting goth user from session", err)
		return nil
	}
	return qorCurrentUser{displayName: gothUser}
}

func (ap ghProvider) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("serveHTTP route", r.URL.Path)

	switch r.URL.Path {
	// login start
	case loginPath:
		u, err := gothic.GetAuthURL(w, r)
		if err != nil {
			fmt.Println("login start error", err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err)
			return
		}
		fmt.Println("login handler redirecting to", u)
		http.Redirect(w, r, u, http.StatusTemporaryRedirect)

	// GH callback
	case callbackPath:
		gothUsr, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			fmt.Println("error completing goth user auth", err)
			http.Error(w, "completing user auth", http.StatusBadRequest)
			return
		}
		if err := gothic.StoreInSession(userIDKey, gothUsr.Name, r, w); err != nil {
			fmt.Println("error storing goth session", err)
			http.Error(w, "storing goth session", http.StatusBadRequest)
			return
		}
		fmt.Fprint(w, "callback finished")
	// logout
	case logoutPath:
		if err := gothic.Logout(w, r); err != nil {
			fmt.Println("error logging out", err)
			http.Error(w, "logging out", http.StatusBadRequest)
			return
		}
		w.Header().Set("Location", "/admin")
		w.WriteHeader(http.StatusTemporaryRedirect)
	default:
		fmt.Println("unhandled path", r.URL.Path)
		http.NotFound(w, r)
	}
}
