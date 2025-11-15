package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

// Returns a redirect to the discord oauth2 flow
func HandleLinkedRolesRedirect(c echo.Context, oauthConfig *oauth2.Config) error {
	// Generate a random state
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	state := hex.EncodeToString(b)

	// Save the state in the session
	sess, _ := session.Get("session", c)
	sess.Values["state"] = state
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, oauthConfig.AuthCodeURL(state))
}

// Handle the callback from the discord oauth2 flow
func HandleLinkedRolesCallback(c echo.Context, oauthConfig *oauth2.Config) error {
	// Get the state from the session
	sess, _ := session.Get("session", c)
	state, ok := sess.Values["state"].(string)
	if !ok {
		return c.String(http.StatusBadRequest, "state not found in session")
	}

	q := c.Request().URL.Query()
	// A safeguard against CSRF attacks.
	if q["state"][0] != state {
		return c.String(http.StatusBadRequest, "invalid state")
	}

	// Fetch the tokens with code we've received.
	tokens, err := oauthConfig.Exchange(c.Request().Context(), q["code"][0])
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Construct a temporary session with user's OAuth2 access_token.
	ts, err := discordgo.New("Bearer " + tokens.AccessToken)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Retrive the user data.
	u, err := ts.User("@me")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	// For now, just return the user as JSON
	return c.JSON(http.StatusOK, u)
}
