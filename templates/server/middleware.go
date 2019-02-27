package main

import (
	"net/http"
	"net/url"
	"path"

	"github.com/labstack/echo"
)

func cors(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
		c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Token, Response-Type")
		c.Response().Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		c.Response().Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
		if c.Request().Method == "OPTIONS" {
			return c.NoContent(http.StatusOK)
		}
		return next(c)
	}
}

func authenticateRequest(c echo.Context) error {
	return nil
}

func auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := authenticateRequest(c); err != nil {
			return c.JSON(http.StatusUnauthorized, "unathorized")
		}
		return next(c)
	}
}

func sendBinaryFiles(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			httpErr, ok := err.(*echo.HTTPError)
			if ok && httpErr.Code == http.StatusNotFound {
				var urlPath string

				{
					u, err := url.Parse(c.Request().URL.Path)
					if err != nil {
						return err
					}
					urlPath = u.String()[1:]
				}
				if data, err := Asset(urlPath); err == nil {
					switch path.Ext(urlPath) {
					case ".html":
						return c.HTMLBlob(http.StatusOK, data)
					case ".css":
						return c.Blob(http.StatusOK, "text/css", data)
					case ".js":
						return c.Blob(http.StatusOK, "text/javascript", data)
					case ".svg":
						return c.Blob(http.StatusOK, "image/svg+xml", data)
					case ".ico":
						return c.Blob(http.StatusOK, "image/ico", data)
					case ".map", "json":
						return c.Blob(http.StatusOK, "application/json", data)
					case ".woff":
						return c.Blob(http.StatusOK, "font/woff", data)
					case ".woff2":
						return c.Blob(http.StatusOK, "font/woff2", data)
					case "tff":
						return c.Blob(http.StatusOK, "application/octet-stream", data)
					default:
						return c.JSON(http.StatusInternalServerError, "unknown content type: "+path.Ext(urlPath))
					}
				}

				if data, err := Asset("index.html"); err != nil {
					return c.JSON(http.StatusInternalServerError, "failed to locate binary")
				} else {
					return c.HTMLBlob(http.StatusOK, data)
				}
			}
		}
		return err
	}
}
