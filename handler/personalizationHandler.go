package handler

import (
	"github.com/labstack/echo"
	"encoding/json"
	"net/http"
)

func GetRecentlyPlayedTracksHandler(c echo.Context) error {

	res, err := Token.GetRecentlyPlayedTracks()
	if err != nil {
		return err
	}
	out, err := json.Marshal(res)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, string(out))
}