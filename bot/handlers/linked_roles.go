package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/h3mmy/bloopyboi/bot/services"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

const oauth_state_key = "oauth_state"

// Returns a redirect to the discord oauth2 flow
func HandleLinkedRolesRedirect(logger *zap.Logger, c echo.Context, oauthConfig *oauth2.Config) error {
	// Generate state for CSRF protection
	state, err := generateState()
	if err != nil {
		logger.Error("failed to generate state", zap.Error(err))
		return c.String(http.StatusInternalServerError, "Failed to generate state")
	}

	// Save the state in the session
	sess, err := session.Get("session", c)
	if err != nil {
		logger.Error("failed to get session", zap.Error(err))
		return c.String(http.StatusInternalServerError, "Failed to get session")
	}
	sess.Values[oauth_state_key] = state
	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		logger.Error("failed to save session", zap.Error(err))
		return c.String(http.StatusInternalServerError, "Failed to save session")
	}

	logger.Debug("redirecting to discord oauth2 flow", zap.String("state", state))
	return c.Redirect(http.StatusMovedPermanently, oauthConfig.AuthCodeURL(state))
}

// Handle the callback from the discord oauth2 flow
func HandleLinkedRolesCallback(logger *zap.Logger, c echo.Context, oauthConfig *oauth2.Config, discordSvc *services.DiscordService) error {
	// Get the state from the session
	sess, err := session.Get("session", c)
	if err != nil {
		logger.Error("failed to get session", zap.Error(err))
		return c.String(http.StatusInternalServerError, "Failed to get session")
	}
	state, ok := sess.Values[oauth_state_key].(string)
	if !ok {
		logger.Error("state not found in session")
		return c.String(http.StatusBadRequest, "state not found in session")
	}

	q := c.Request().URL.Query()
	// A safeguard against CSRF attacks.
	if len(q["state"]) == 0 {
		logger.Error("state not found in query")
		return c.String(http.StatusBadRequest, "state not found in query")
	}
	if q["state"][0] != state {
		logger.Error("invalid state", zap.String("expected", state), zap.String("got", q["state"][0]))
		return c.String(http.StatusBadRequest, "invalid state")
	}

	// Fetch the tokens with code we've received.
	tokens, err := oauthConfig.Exchange(c.Request().Context(), q["code"][0])
	if err != nil {
		logger.Error("failed to exchange token", zap.Error(err))
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Construct a temporary session with user's OAuth2 access_token.
	ts, err := discordgo.New("Bearer " + tokens.AccessToken)
	if err != nil {
		logger.Error("failed to create discord session", zap.Error(err))
		return c.String(http.StatusInternalServerError, err.Error())
	}

	rcData, err := discordSvc.UpdateDiscordUserRoleConnection(c.Request().Context(), ts, oauthConfig)
	if err != nil {
		logger.Error("failed to update discord user role connection", zap.Error(err))
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
