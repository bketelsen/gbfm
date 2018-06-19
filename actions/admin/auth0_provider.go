package admin

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
	"github.com/qor/admin"
	"github.com/qor/auth"
	"github.com/qor/auth/claims"
	"github.com/qor/qor"
)

// auth0Provider is a Provider implementation that uses Auth0
// it also implements admin.Auth
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

// /login/github
func (ap auth0Provider) Login(c *auth.Context) {
	fmt.Fprintln(c.Writer, "auth0 login")
	// fmt.Println("auth0 login")
	// gothic.BeginAuthHandler(c.Writer, c.Request)
}

// /logout/github
func (ap auth0Provider) Logout(c *auth.Context) {
	fmt.Fprintln(c.Writer, "auth0 logout")
	// w, r := c.Writer, c.Request
	// gothic.Logout(w, r)
	// w.Header().Set("Location", "/")
	// w.WriteHeader(http.StatusTemporaryRedirect)
}

func (ap auth0Provider) Register(c *auth.Context) {
	fmt.Fprintln(c.Writer, "auth0 register")
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
	fmt.Fprintln(c.Writer, "auth0 callback")
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

func (ap auth0Provider) addAuthRoutes(httpMux *http.ServeMux) {
	router := mux.NewRouter()

	// GH callback
	router.HandleFunc("/auth/{provider}/callback", func(w http.ResponseWriter, r *http.Request) {
		_, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			fmt.Fprintln(w, r)
			return
		}
		// t, _ := template.New("foo").Parse(userTemplate)
		// t.Execute(res, user)
	}).Methods("GET")

	// logout
	router.HandleFunc("/logout/{provider}", func(w http.ResponseWriter, r *http.Request) {
		gothic.Logout(w, r)
		w.Header().Set("Location", "/")
		w.WriteHeader(http.StatusTemporaryRedirect)
	}).Methods("GET")

	// login start
	router.HandleFunc("/auth/{provider}", func(w http.ResponseWriter, r *http.Request) {
		if _, err := gothic.CompleteUserAuth(w, r); err == nil {
			// t, _ := template.New("foo").Parse(userTemplate)
			// t.Execute(res, gothUser)
		} else {
			gothic.BeginAuthHandler(w, r)
		}
	})
	httpMux.Handle("/auth", router)
}

func (ap auth0Provider) ServeHTTP(c *auth.Context) {}

func (ap auth0Provider) LoginURL(c *admin.Context) string {
	return "/auth/github"
}

func (ap auth0Provider) LogoutURL(c *admin.Context) string {
	return "/logout/github"
}

func (ap auth0Provider) GetCurrentUser(c *admin.Context) qor.CurrentUser {
	gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		return nil
	}

	displayName := gothUser.Email
	if displayName == "" {
		displayName = gothUser.Name
	}
	if displayName == "" {
		displayName = gothUser.UserID
	}
	return qorCurrentUser{displayName: displayName}
}
