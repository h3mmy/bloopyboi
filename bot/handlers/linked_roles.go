package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/h3mmy/bloopyboi/bot/services"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

const oauth_state_key = "oauth_state"

// Returns a redirect to the discord oauth2 flow
func HandleLinkedRolesRedirect(c echo.Context, oauthConfig *oauth2.Config) error {

	// Generate state for CSRF protection
	state, err := generateState()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to generate state")
	}

	// Save the state in the session
	sess, _ := session.Get("session", c)
	sess.Values[oauth_state_key] = state
	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to save session")
	}

	return c.Redirect(http.StatusMovedPermanently, oauthConfig.AuthCodeURL(state))
}

// Handle the callback from the discord oauth2 flow
func HandleLinkedRolesCallback(c echo.Context, oauthConfig *oauth2.Config, discordSvc *services.DiscordService) error {
	// Get the state from the session
	sess, _ := session.Get("session", c)
	state, ok := sess.Values[oauth_state_key].(string)
	if !ok {
		return c.String(http.StatusBadRequest, "state not found in session")
	}

	q := c.Request().URL.Query()
	// A safeguard against CSRF attacks.
	if len(q["state"]) == 0 {
		return c.String(http.StatusBadRequest, "state not found in query")
	}
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

	rcData, err := discordSvc.UpdateDiscordUserRoleConnection(c.Request().Context(), ts, oauthConfig)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	// For now, just return the user as JSON
	return c.JSON(http.StatusOK, rcData)
}

// generateState creates a cryptographically secure random string for OAuth2 'state'
func generateState() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
