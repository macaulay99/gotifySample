/**
Don't Edit
*/

package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/gericass/gotify"
)

var Auth *gotify.Client

func ExampleHandler(c echo.Context) error {
	url := Auth.AuthURL()
	return c.Redirect(301, url)
}

func ExampleCallbackHandler(c echo.Context) error {
	token, err := Auth.Token(c.Request())
	if err != nil {
		return err
	}

	res, err := gotify.GetUserStatus(token)
	if err != nil {
		return nil
	}
	return c.String(http.StatusOK, res)
}
