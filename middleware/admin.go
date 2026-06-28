package middleware

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
)

func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {

		role, ok := c.Get("user_role").(string)

		fmt.Printf("DEBUG ROLE: '%v' OK: %v\n", role, ok)

		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]any{
				"success": false,
				"message": "User role missing from context",
			})
		}

		if role != "admin" {
			return c.JSON(http.StatusForbidden, map[string]any{
				"success": false,
				"message": "Forbidden: Admin only",
			})
		}

		return next(c)
	}
}
