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
	e.GET("/artists/albums", handler.ArtistsAlbumsHandler)

	e.GET("/artists/toptracks", handler.ArtistsTopTracksHandler)
	e.GET("/artists/relatedartists", handler.ArtistsRelatedArtistsHandler)

	e.GET("/browse/featuredplaylists", handler.BrowseFeaturedPlaylists)
	e.GET("/browse/newreleases", handler.BrowseNewReleases)
	e.GET("/browse/categories", handler.BrowseCategories)
	e.GET("/browse/category", handler.BrowseCategory)
	e.GET("/browse/categorysplaylists", handler.BrowseCategorysPlaylists)

	e.GET("/recommendations", handler.RecommendationsHandler)

	e.GET("/following/artists", handler.FollowingArtistsHandler)
	e.GET("/following/follow/artistsorusers", handler.FollowArtistsOrUsersHandler)
	e.GET("/following/unfollow/artistsorusers", handler.UnfollowArtistsOrUsersHandler)


	// Require SSL
	e.Logger.Fatal(e.StartTLS(":3000", "cert.pem", "key.pem"))
}
