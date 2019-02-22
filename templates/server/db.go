package main

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
)

const dbName = "test.db"

// ProvideConnection attaches a gorm db to the gin context
func provideConnection(next echo.HandlerFunc) echo.HandlerFunc {
	// define your own database here
	db := ""
	return func(c echo.Context) error {
		c.Set("db", db)
		return next(c)
	}
}
