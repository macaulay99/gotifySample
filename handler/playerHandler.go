package handler

import (
	"encoding/json"
	"github.com/labstack/echo"
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

	deviceIDs := []string{""}

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

func StartUsersTrackPlayHandler(c echo.Context) error {

	trackinfo := []string{"spotify:track:2tPDBuHnkBNvXNn7NtqtqX",
		"spotify:track:1301WleyT98MSxVHPZCA6M"}
	seekms := 50000
	err := Token.StartUsersTrackPlay(trackinfo, seekms)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "OK")
}

func PauseUsersPlaybackHandler(c echo.Context) error {

	err := Token.PauseUsersPlayback()
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "OK")
}

func SkipUsersPlaybackToNextHandler(c echo.Context) error {

	err := Token.SkipUsersPlaybackToNext()
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "OK")
}

func SkipUsersPlaybackToPreviousHandler(c echo.Context) error {

	err := Token.SkipUsersPlaybackToPrevious()
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "OK")
}

func SeekToPositionInCurrentlyPlayingTrackHandler(c echo.Context) error {

	pos := 60000

	print("SeekToPositionInCurrentlyPlayingTrackHandler", pos)
	err := Token.SeekToPositionInCurrentlyPlayingTrack(pos)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "OK")
}

func SetRepeatModeUsersPlaybackHandler(c echo.Context) error {

	err := Token.SetRepeatModeUsersPlayback("off")
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "OK")
}

func SetVolumeUsersPlaybackHandler(c echo.Context) error {

	err := Token.SetVolumeUsersPlayback(20)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "OK")
}

func ToggleShuffleUsersPlaybackHandler(c echo.Context) error {

	err := Token.ToggleShuffleUsersPlayback(false)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "OK")
}
