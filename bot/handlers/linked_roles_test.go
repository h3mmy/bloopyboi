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
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

func TestHandleLinkedRolesRedirect(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	oauthConfig := &oauth2.Config{}
	logger, _ := zap.NewDevelopment()

	e.GET("/", func(c echo.Context) error {
		return HandleLinkedRolesRedirect(logger, c, oauthConfig)
	})

	e.ServeHTTP(rec, req)

	// Assertions
	assert.Equal(t, http.StatusMovedPermanently, rec.Code)
	// Check that a session cookie was set
	s := &http.Cookie{}
	s.Name = "session"
	for _, cookie := range rec.Result().Cookies() {
		if cookie.Name == "session" {
			s = cookie
		}
	}
	assert.NotEmpty(t, s.Value, "session cookie should be set")

	// check that the session contains a non-nil state
	store := sessions.NewCookieStore([]byte("secret"))
	session, err := store.Get(req, "session")
	assert.NoError(t, err)
	assert.NotEmpty(t, session.Values[oauth_state_key], "state should be set in session")

}

func TestHandleLinkedRolesCallback_NoStateInQuery(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	oauthConfig := &oauth2.Config{}
	discordSvc := &services.DiscordService{}
	logger, _ := zap.NewDevelopment()

	e.GET("/", func(c echo.Context) error {
		// Stub the session
		sess, _ := session.Get("session", c)
		sess.Values[oauth_state_key] = "test_state"
		err := sess.Save(c.Request(), c.Response())
		assert.NoError(t, err)
		return HandleLinkedRolesCallback(logger, c, oauthConfig, discordSvc)
	})

	e.ServeHTTP(rec, req)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "state not found in query")
}
