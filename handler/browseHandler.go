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