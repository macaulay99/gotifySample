package handler

import (
	"github.com/labstack/echo"
	"encoding/json"
	"net/http"
)

func BrowseFeaturedPlaylists(c echo.Context) error {

	res, err := Token.GetBrowseFeaturedPlaylists()
	if err != nil {
		return nil
	}
	out, err := json.Marshal(res)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, string(out))
}

func BrowseNewReleases(c echo.Context) error {

	res, err := Token.GetBrowseNewReleases()
	if err != nil {
		return nil
	}
	out, err := json.Marshal(res)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, string(out))
}

func BrowseCategories(c echo.Context) error {

	res, err := Token.GetBrowseCategories()
	if err != nil {
		return nil
	}
	out, err := json.Marshal(res)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, string(out))
}

func BrowseCategory(c echo.Context) error {

	res, err := Token.GetBrowseCategory("party")
	if err != nil {
		return nil
	}
	out, err := json.Marshal(res)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, string(out))
}
