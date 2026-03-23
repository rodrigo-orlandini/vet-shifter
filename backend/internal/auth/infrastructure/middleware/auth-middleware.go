package middleware

import (
	"net/http"
	"strings"

	"rodrigoorlandini/vet-shifter/internal/_shared/utils"

	"github.com/gin-gonic/gin"
)

const AuthUserIDKey = "auth_user_id"
const AuthUserTypeKey = "auth_user_type"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		var tokenString string

		if authHeader == "" {
			if cookieToken, err := c.Cookie(AccessTokenCookieName); err == nil {
				tokenString = cookieToken
			}
		} else {
			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || parts[0] != "Bearer" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"code":  "UNAUTHORIZED",
					"error": "invalid or missing token",
				})
				return
			}

			tokenString = parts[1]
		}

		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":  "UNAUTHORIZED",
				"error": "invalid or missing token",
			})
			return
		}

		claims, err := utils.VerifyJWT(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":  "UNAUTHORIZED",
				"error": "invalid or missing token",
			})
			return
		}

		c.Set(AuthUserIDKey, claims.Sub)
		c.Set(AuthUserTypeKey, claims.Type)
		c.Next()
	}
}
