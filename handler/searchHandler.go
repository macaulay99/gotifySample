package handler

import (
	"github.com/labstack/echo"
	"net/http"
	"encoding/json"
)

func SearchArtistsHandler(c echo.Context) error {

	res, err := Token.SearchArtists("burial")
	if err != nil {
		return err
	}
	out, err := json.Marshal(res)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, string(out))
}

func SearchAlbumsHandler(c echo.Context) error {

	res, err := Token.SearchAlbums("untrue")
	if err != nil {
		return err
	}
	out, err := json.Marshal(res)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, string(out))
}
