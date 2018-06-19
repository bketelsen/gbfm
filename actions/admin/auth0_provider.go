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

// auth0Provider is a Provider implementation that uses Auth0
// it also implements admin.Auth.
//
// don't create this directly, use newAuth0Provider instead. it needs
// to set up global state
type auth0Provider struct {
	provider goth.Provider
}

// TODO: rename this to GH provider
func newAuth0Provider(ghKey, ghSecret, host string) providerAuther {
	ghProvider := github.New(ghKey, ghSecret, fmt.Sprintf("%s%s", host, callbackPath))
	goth.UseProviders(ghProvider)
	return &auth0Provider{provider: ghProvider}
}

func (ap auth0Provider) GetName() string {
	return "auth0"
}

func (ap auth0Provider) ConfigAuth(a *auth.Auth) {
	a.Config.LoginHandler = func(c *auth.Context, fn func(*auth.Context) (*claims.Claims, error)) {

	}
	a.Config.LogoutHandler = ap.Logout
}

// /login/github
func (ap auth0Provider) Login(c *auth.Context) {
	// fmt.Println("auth0 login")
	// now := time.Now()
	// c.Auth.Login(c.Writer, c.Request, &claims.Claims{
	// 	Provider:    ap.GetName(),
	// 	UserID:      "aaron",
	// 	LastLoginAt: &now,
	// })
	// fmt.Fprint(c.Writer, "logged in!")
	fmt.Println("login!")
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

// /logout/github
func (ap auth0Provider) Logout(c *auth.Context) {
	fmt.Println("logout!")
	// w, r := c.Writer, c.Request
	// gothic.Logout(w, r)
	// w.Header().Set("Location", "/")
	// w.WriteHeader(http.StatusTemporaryRedirect)
}

func (ap auth0Provider) Register(c *auth.Context) {
	fmt.Println("register!")
	w, r := c.Writer, c.Request

	// try to see if the person is logged in already
	gothUser, err := gothic.CompleteUserAuth(w, r)
	if err == nil {
		c.SessionStorerInterface.Update(w, r, &claims.Claims{
			Provider: ap.GetName(),
			UserID:   gothUser.UserID,
		})
		return
	}

	// otherwise start auth
	gothic.BeginAuthHandler(w, r)
	return

}

func (ap auth0Provider) Callback(c *auth.Context) {
	fmt.Println("callback!")
	w, r := c.Writer, c.Request
	gothUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, "auth failed!")
		return
	}
	c.SessionStorer.Update(w, r, &claims.Claims{
		Provider: ap.GetName(),
		UserID:   gothUser.UserID,
	})
	return
}

func (ap auth0Provider) ServeHTTP(c *auth.Context) {
	fmt.Println("serve http!")
	router := mux.NewRouter()

	// login start
	router.HandleFunc(loginPath, func(w http.ResponseWriter, r *http.Request) {
		// if _, err := gothic.CompleteUserAuth(w, r); err == nil {
		// 	// t, _ := template.New("foo").Parse(userTemplate)
		// 	// t.Execute(res, gothUser)
		// } else {
		r.URL.Query().Set("provider", "github")
		gothic.BeginAuthHandler(w, r)
		// }
	})

	// GH callback
	router.HandleFunc(callbackPath, func(w http.ResponseWriter, r *http.Request) {
		_, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			fmt.Fprintln(w, r)
			return
		}
		// t, _ := template.New("foo").Parse(userTemplate)
		// t.Execute(res, user)
	}).Methods("GET")

	// logout
	router.HandleFunc(logoutPath, func(w http.ResponseWriter, r *http.Request) {
		gothic.Logout(w, r)
		w.Header().Set("Location", "/")
		w.WriteHeader(http.StatusTemporaryRedirect)
	}).Methods("GET")

	router.ServeHTTP(c.Writer, c.Request)
}

func (ap auth0Provider) LoginURL(c *admin.Context) string {
	fmt.Println("login url!")
	return "/admin/auth/login?provider=github"
}

func (ap auth0Provider) LogoutURL(c *admin.Context) string {
	fmt.Println("logout url!")
	return "/admin/auth/logout?provider=github"
}

func (ap auth0Provider) GetCurrentUser(c *admin.Context) qor.CurrentUser {
	fmt.Println("get current user!")
	gothUser, err := gothic.GetFromSession(userIDKey, c.Request)
	if err != nil {
		gothic.BeginAuthHandler(c.Writer, c.Request)
		return nil
	}
	return qorCurrentUser{displayName: gothUser}
	// gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	// if err != nil {
	// 	return nil
	// }

	// displayName := gothUser.Email
	// if displayName == "" {
	// 	displayName = gothUser.Name
	// }
	// if displayName == "" {
	// 	displayName = gothUser.UserID
	// }
	// return qorCurrentUser{displayName: displayName}
}
