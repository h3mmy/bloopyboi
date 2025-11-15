package handlers

import (
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

// Returns a redirect to the discord oauth2 flow
func HandleLinkedRolesRedirect(c echo.Context, oauthConfig *oauth2.Config) error {
	return c.Redirect(http.StatusMovedPermanently, oauthConfig.AuthCodeURL("random-state"))
}

// Handle the callback from the discord oauth2 flow
func HandleLinkedRolesCallback(c echo.Context, oauthConfig *oauth2.Config) error {
	q := c.Request().URL.Query()
	// A safeguard against CSRF attacks.
	// Usually tied to requesting user or random.
	// NOTE: Hardcoded for the sake of the example.
	if q["state"][0] != "random-state" {
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
