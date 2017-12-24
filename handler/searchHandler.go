package handler

import (
	"github.com/labstack/echo"
	"net/http"
	"encoding/json"
)

func SearchArtistHandler(c echo.Context) error {

	res, err := Token.SearchArtist("burial")
	if err != nil {
		return err
	}
	out, err := json.Marshal(res)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, string(out))
}
