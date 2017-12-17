package handler

import (
	"github.com/labstack/echo"
	"encoding/json"
	"net/http"
)

func FollowingArtistsHandler(c echo.Context) error {

	res, err := Token.GetFollowingArtists()
	if err != nil {
		return err
	}
	out, err := json.Marshal(res)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, string(out))
}

func FollowArtistsOrUsersHandler(c echo.Context) error {

	ids := []string{"0uCCBpmg6MrPb1KY2msceF"}
	err := Token.FollowArtistsOrUsers("artist", ids)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "ok")
}
