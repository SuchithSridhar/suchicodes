package cookie

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/suchithsridhar/suchicodes/configs"
)

const UserKey = "username"
const ThemeKey = "theme"

const dayInSeconds = 24 * 60 * 60
const yearInSeconds = 365 * dayInSeconds
const defaultTheme = "light"
const defaultUser = ""

func SetThemeCookie(theme string, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = ThemeKey
	cookie.Value = theme
	cookie.MaxAge = yearInSeconds
	c.SetCookie(cookie)
}

func ReadThemeCookie(c echo.Context) string {
	cookie, err := c.Cookie(ThemeKey)
	if err != nil {
		return defaultTheme
	}
	return cookie.Value
}

func ReadUserSession(c echo.Context) (username string, ok bool) {
	sesh, err := session.Get("session", c)
	if err != nil {
		return "", false
	}

	if user, ok := sesh.Values[UserKey]; !ok {
		return "", false
	} else if userStr, ok := user.(string); !ok {
		return "", false
	} else {
		return userStr, true
	}
}

func SetUserSession(username string, c echo.Context) (ok bool) {
	sesh, err := session.Get("session", c)
	if err != nil {
		slog.Error("Unable to get session during login.")
		return false
	}

	sesh.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   dayInSeconds,
		HttpOnly: true,
		// only when env is not dev
		Secure: configs.Config.ENVIRONMENT != "development",
	}

	sesh.Values[UserKey] = username
	sesh.Save(c.Request(), c.Response())

	return true
}

func DeleteUserSession(c echo.Context) (ok bool) {
	sesh, err := session.Get("session", c)
	if err != nil {
		slog.Error("Unable to get session during logout.")
		return false
	}

	delete(sesh.Values, UserKey)
	sesh.Save(c.Request(), c.Response())

	return true
}
