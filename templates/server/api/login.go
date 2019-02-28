package api

import (
	"errors"
	"net/http"

	"{{.ServerAPIPath}}/service"
	"github.com/labstack/echo"
)

type loginRequst struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l *loginRequst) Validate() error {
	if l.Username == "" {
		return errors.New("username cannot be blank")
	}
	if l.Password == "" {
		return errors.New("password cannot be blank")
	}
	// hit auth service
	return nil
}

func loginHandler(c echo.Context) error {
	login := &loginRequst{}
	if err := c.Bind(login); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := login.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	jwt, err := service.NewJWT(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	token, err := jwt.NewToken(login.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, token)
}
