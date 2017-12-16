package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/gericass/gotify"
)

var Auth *gotify.Client
var Token *gotify.Tokens

func Handler(c echo.Context) error {
	url := Auth.AuthURL()
	return c.Redirect(301, url)
}

func CallbackHandler(c echo.Context) error {
	t, err := Auth.GetToken(c.Request())
	if err != nil {
		return err
	}
	Token = t

	return c.String(http.StatusOK, "Authentication success")
}
