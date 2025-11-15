package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/h3mmy/bloopyboi/bot/services"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

func TestHandleLinkedRolesCallback_NoStateInQuery(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	oauthConfig := &oauth2.Config{}
	discordSvc := &services.DiscordService{}

	e.GET("/", func(c echo.Context) error {
		// Stub the session
		sess, _ := session.Get("session", c)
		sess.Values[oauth_state_key] = "test_state"
		sess.Save(c.Request(), c.Response())
		return HandleLinkedRolesCallback(c, oauthConfig, discordSvc)
	})

	e.ServeHTTP(rec, req)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "state not found in query")
}
