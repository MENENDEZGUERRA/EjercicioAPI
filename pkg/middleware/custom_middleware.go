package middleware

import "github.com/labstack/echo/v4"

func CustomHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("X-Custom-Header", "MyValue")
		return next(c)
	}
}
