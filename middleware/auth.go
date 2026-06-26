package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/noboKumar/SpotSync-server/utils"
)

func AuthMiddleware(jwtService utils.JwtService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "missing authorization header",
				})
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "invalid authorization header format",
				})
			}

			claims, err := jwtService.ValidateToken(parts[1])
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "invalid or expired token",
				})
			}

			// reject refresh tokens from being used as access tokens
			if claims.TokenType != utils.TokenTypeAccess {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "invalid token type",
				})
			}

			c.Set("user_id", claims.UserId)
			c.Set("user_email", claims.Email)
			c.Set("user_name", claims.Name)
			c.Set("user_role", claims.Role)

			return next(c)
		}
	}
}
