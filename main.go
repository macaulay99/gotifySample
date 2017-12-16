/**
Don't Edit
*/

package main

import (
	"github.com/gericass/gotify"
	"github.com/gericass/gotifySample/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const clientID = ""
const clientSecret = ""
const callbackURI = "https://localhost:3000/callback/"

func main() {
	e := echo.New()
	e.Pre(middleware.HTTPSRedirect())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	handler.Auth = gotify.Set(clientID, clientSecret, callbackURI)

	e.GET("/", handler.Handler)
	e.GET("/callback/", handler.CallbackHandler)
	e.GET("/albums", handler.AlbumsHandler)
	e.GET("/albumstracks", handler.AlbumsTracksHandler)
	e.GET("/artists", handler.ArtistsHandler)
	e.GET("/artistsalbums", handler.ArtistsAlbumsHandler)
	e.GET("/artiststoptracks", handler.ArtistsTopTracksHandler)
	e.GET("/artistsrelatedartists", handler.ArtistsRelatedArtistsHandler)


	// Require SSL
	e.Logger.Fatal(e.StartTLS(":3000", "cert.pem", "key.pem"))
}
