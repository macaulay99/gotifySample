package handler

import (
	"github.com/labstack/echo"
	"net/http"
	"encoding/json"
)

func AlbumsHandler(c echo.Context) error {

	albumIDs := []string{"4sgYpkIASM1jVlNC8Wp9oF", "1c9Sx7XdXuMptGyfCB6hHs"}

	res, err := Token.GetAlbums(albumIDs)
	if err != nil {
		return nil
	}
	out, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	return c.String(http.StatusOK, string(out))
}
