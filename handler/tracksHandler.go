package handler

import (
	"github.com/labstack/echo"
	"encoding/json"
	"net/http"
)

func GetTracksHandler(c echo.Context) error {

	tracks := []string{"3n3Ppam7vgaVa1iaRUc9Lp"}

	res, err := Token.GetTracks(tracks)
	if err != nil {
		return err
	}
	out, err := json.Marshal(res)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, string(out))
}


func GetAudioAnalysisHandler(c echo.Context) error {

	track := "3n3Ppam7vgaVa1iaRUc9Lp"

	res, err := Token.GetAudioAnalysis(track)
	if err != nil {
		return err
	}
	out, err := json.Marshal(res)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, string(out))
}
