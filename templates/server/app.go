package main

import (
	"net/http"

	"{{.ServerAPIPath}}"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewApp() *echo.Echo {
	engine := echo.New()
	engine.Debug = true
	engine.Use(middleware.Recover())
	engine.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[ECHO] - ${time_rfc3339} |${status}| ${latency_human} | ${host} | ${method} ${uri}\n",
	}))

	engine.Use(cors)

	engine.GET("/", func(c echo.Context) error {
		if data, err := Asset("index.html"); err == nil {
			return c.HTMLBlob(http.StatusOK, data)
		} else {
			return err
		}
	})

	api.Routes(engine.Group("/api"))

	engine.Use(sendBinaryFiles)

	return engine
}
