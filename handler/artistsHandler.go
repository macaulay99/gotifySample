package handler

import (
	"github.com/labstack/echo"
	"encoding/json"
	"net/http"
)

func ArtistsHandler(c echo.Context) error {

	artistIDs := []string{"74XFHRwlV6OrjEM0A2NCMF", "0uCCBpmg6MrPb1KY2msceF"}

	res, err := Token.GetArtists(artistIDs)
	if err != nil {
		return nil
	}
	out, err := json.Marshal(res)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, string(out))
}

func ArtistsAlbumsHandler(c echo.Context) error {

	artistID := "74XFHRwlV6OrjEM0A2NCMF"

	res, err := Token.GetArtistsAlbums(artistID)
	if err != nil {
		return nil
	}
	out, err := json.Marshal(res)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, string(out))
}

func ArtistsTopTracksHandler(c echo.Context) error {

	artistID := "74XFHRwlV6OrjEM0A2NCMF"
	country := "JP"

	res, err := Token.GetArtistsTopTracks(artistID,country)
	if err != nil {
		return nil
	}
	out, err := json.Marshal(res)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, string(out))
}