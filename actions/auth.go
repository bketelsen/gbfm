package actions

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/murlokswarm/log"
	"github.com/pkg/errors"

	"github.com/gophersnacks/gbfm/models"
	"golang.org/x/crypto/bcrypt"
)

// AuthNew loads the signin page
func AuthNew(c buffalo.Context) error {
	c.Set("user", models.User{})
	return c.Render(200, r.HTML("auth/new.html"))
}

// LoginRequest represents a login form.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AuthToken attempts to log the user in with an existing account
// and returns a JWT token with success
func AuthToken(c buffalo.Context) error {
	decoder := json.NewDecoder(c.Request().Body)
	var l LoginRequest
	err := decoder.Decode(&l)
	if err != nil {
		fmt.Println(l)
		return err
	}
	fmt.Println(l)
	tx := c.Value("tx").(*pop.Connection)

	u := &models.User{}
	// find a user with the email
	err = tx.Where("email = ?", strings.ToLower(l.Email)).First(u)

	// helper function to handle bad attempts
	bad := func() error {
		c.Set("user", u)
		verrs := validate.NewErrors()
		verrs.Add("email", "invalid email/password")
		c.Set("errors", verrs)
		return c.Error(http.StatusUnauthorized, verrs)
	}
	// Todo: return better description, probably in JSON?
	if err != nil {
		fmt.Println("Error:", err)
		if errors.Cause(err) == sql.ErrNoRows {
			// couldn't find an user with the supplied email address.
			return bad()
		}
		return errors.WithStack(err)
	}

	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(oneWeek()).Unix(),
		Issuer:    fmt.Sprintf("%s.gophersnacks.com", envy.Get("GO_ENV", "development")),
		Id:        string(u.ID),
	}

	signingKey, err := ioutil.ReadFile(envy.Get("JWT_KEY_PATH", ""))

	if err != nil {
		return fmt.Errorf("could not open jwt key, %v", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(signingKey)

	if err != nil {
		return fmt.Errorf("could not sign token, %v", err)
	}

	return c.Render(200, r.JSON(map[string]string{"token": tokenString}))
}

// AuthCreate attempts to log the user in with an existing account.
func AuthCreate(c buffalo.Context) error {
	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return errors.WithStack(err)
	}

	tx := c.Value("tx").(*pop.Connection)

	// find a user with the email
	err := tx.Where("email = ?", strings.ToLower(u.Email)).First(u)

	// helper function to handle bad attempts
	bad := func() error {
		c.Set("user", u)
		verrs := validate.NewErrors()
		verrs.Add("email", "invalid email/password")
		c.Set("errors", verrs)
		return c.Render(422, r.HTML("auth/new.html"))
	}

	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			// couldn't find an user with the supplied email address.
			return bad()
		}
		return errors.WithStack(err)
	}

	// confirm that the given password matches the hashed password from the db
	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(u.PasswordHash))
	if err != nil {
		return bad()
	}
	c.Session().Set("current_user_id", u.ID)
	c.Flash().Add("success", "Welcome Back to Buffalo!")

	return c.Redirect(302, "/")
}

// AuthDestroy clears the session and logs a user out
func AuthDestroy(c buffalo.Context) error {
	c.Session().Clear()
	c.Flash().Add("success", "You have been logged out!")
	return c.Redirect(302, "/")
}

func oneWeek() time.Duration {
	return 7 * 24 * time.Hour
}

// RestrictedHandlerMiddleware searches and parses the jwt token in order to authenticate
// the request and populate the Context with the user contained in the claims.
func RestrictedHandlerMiddleware(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")

		if len(tokenString) == 0 {
			return c.Error(http.StatusUnauthorized, fmt.Errorf("No token set in headers"))
		}

		// parsing token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// key
			mySignedKey, err := ioutil.ReadFile(envy.Get("JWT_KEY_PATH", ""))

			if err != nil {
				return nil, fmt.Errorf("could not open jwt key, %v", err)
			}

			return mySignedKey, nil
		})

		if err != nil {
			return c.Error(http.StatusUnauthorized, fmt.Errorf("Could not parse the token, %v", err))
		}

		// getting claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			log.Errorf("claims: %v", claims)

			tx := c.Value("tx").(*pop.Connection)
			u := &models.User{}
			// find a user with the email
			err := tx.Find(&u, claims["jti"].(string))
			if err != nil {
				return c.Error(http.StatusUnauthorized, fmt.Errorf("Could not identify the user"))
			}

			c.Set("user", u)
			c.Session().Set("current_user_id", u.ID)

		} else {
			return c.Error(http.StatusUnauthorized, fmt.Errorf("Failed to validate token: %v", claims))
		}

		return next(c)
	}
}
