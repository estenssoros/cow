package api

import (
	"github.com/labstack/echo"
)

func Routes(e *echo.Group) {
	e.POST("/login", loginHandler)
}
