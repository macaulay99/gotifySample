package handler

import (
	"github.com/labstack/echo"
	"encoding/json"
	"net/http"
)

func GetUsersAvailableDevicesHandler(c echo.Context) error {

	res, err := Token.GetUsersAvailableDevices()
	if err != nil {
		return err
	}
	out, err := json.Marshal(res)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, string(out))
}

/*
func GetInformationAboutUsersCurrentPlaybackHandler(c echo.Context) error {

	res, err := Token.GetInformationAboutUsersCurrentPlayback()
	if err != nil {
		return err
	}
	out, err := json.Marshal(res)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, string(out))
}
*/

/*
func GetUsersCurrentlyPlayingTrackHandler(c echo.Context) error {

	res, err := Token.GetUsersCurrentlyPlayingTrack()
	if err != nil {
		return err
	}
	out, err := json.Marshal(res)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, string(out))
}
*/

func TransferUsersPlaybackHandler(c echo.Context) error {

	deviceIDs := []string{"74ASZWbe4lXaubB36ztrGX"}

	err := Token.TransferUsersPlayback(deviceIDs)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "OK")
}

func StartResumeUsersPlaybackHandler(c echo.Context) error {


	err := Token.StartResumeUsersPlayback()
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "OK")
}