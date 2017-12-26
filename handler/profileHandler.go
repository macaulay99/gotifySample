package handler

import (
	"github.com/labstack/echo"
	"encoding/json"
	"net/http"
)

func GetCurrentUsersProfileHandler(c echo.Context) error {

	res, err := Token.GetCurrentUsersProfile()
	if err != nil {
		return err
	}
	out, err := json.Marshal(res)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, string(out))
}

func GetUsersProfileHandler(c echo.Context) error {

	res, err := Token.GetUsersProfile("tuggareutangranser")
	if err != nil {
		return err
	}
	out, err := json.Marshal(res)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, string(out))
}