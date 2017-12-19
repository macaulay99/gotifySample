package handler

import (
	"github.com/labstack/echo"
	"net/http"
	"encoding/json"
)

func SaveTracksHandler(c echo.Context) error {

	ids := []string{""}

	err := Token.SaveTracks(ids)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "OK")
}


func UsersSavedTracksHandler(c echo.Context) error {


	res, err := Token.GetUsersSavedTracks()
	if err != nil {
		return err
	}
	out, err := json.Marshal(res)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, string(out))
}