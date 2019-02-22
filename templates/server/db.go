package main

import (
	"github.com/labstack/echo"
)

// ProvideConnection attaches a gorm db to the gin context
func provideConnection(next echo.HandlerFunc) echo.HandlerFunc {
	// define your own database here
	db := ""
	return func(c echo.Context) error {
		c.Set("db", db)
		return next(c)
	}
}
