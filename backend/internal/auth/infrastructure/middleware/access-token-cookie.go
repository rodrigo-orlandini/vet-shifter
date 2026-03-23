package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

const AccessTokenCookieName = "access_token"

func cookieSecure(c *gin.Context) bool {
	if v := os.Getenv("COOKIE_SECURE"); v != "" {
		return v == "true"
	}

	return c.Request.TLS != nil
}

func SetAccessTokenCookie(c *gin.Context, token string, expiresAt string) {
	exp, err := time.Parse(time.RFC3339, expiresAt)
	if err != nil {
		exp = time.Now().Add(24 * time.Hour)
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     AccessTokenCookieName,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   cookieSecure(c),
		SameSite: http.SameSiteLaxMode,
		Expires:  exp,
	})
}

func ClearAccessTokenCookie(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     AccessTokenCookieName,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   cookieSecure(c),
		SameSite: http.SameSiteLaxMode,
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
	})
}
