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
	e.GET("/refresh/", handler.RefreshHandler)

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
	e.GET("/following/current/artistsorusers", handler.CurrentFollowArtistsOrUsersHandler)
	e.GET("/following/follow/playlists", handler.FollowPlaylistHandler)
	e.GET("/following/unfollow/playlists", handler.UnfollowPlaylistHandler)
	e.GET("/following/check/playlist", handler.CheckFollowPlaylistHandler)

	e.GET("/library/save/tracks", handler.SaveTracksHandler)
	e.GET("/library/get/tracks", handler.GetUsersSavedTracksHandler)
	e.GET("/library/delete/tracks", handler.RemoveUsersSavedTracksHandler)
	e.GET("/library/check/tracks", handler.CheckUsersSavedTracksHandler)
	e.GET("/library/save/albums", handler.SaveAlbumsHandler)
	e.GET("/library/get/albums", handler.GetUsersSavedAlbumsHandler)
	e.GET("/library/remove/albums", handler.RemoveAlbumsForCurrentUserHandler)
	e.GET("/library/check/albums", handler.CheckUsersSavedAlbumsHandler)

	e.GET("/personalization/get/recentlyplayedtracks", handler.GetRecentlyPlayedTracksHandler)

	e.GET("/player/get/availabledevices", handler.GetUsersAvailableDevicesHandler)

	// e.GET("/player/get/informationaboutuserscurrentplayback", handler.GetInformationAboutUsersCurrentPlaybackHandler)
	// e.GET("/player/get/currentlyplaying", handler.GetUsersCurrentlyPlayingTrackHandler)
	e.GET("/player/put/usersplayback", handler.TransferUsersPlaybackHandler)
	e.GET("/player/put/startresume", handler.StartResumeUsersPlaybackHandler)
	e.GET("/player/put/pause", handler.PauseUsersPlaybackHandler)
	e.GET("/player/post/skipnext", handler.SkipUsersPlaybackToNextHandler)
	e.GET("/player/post/skipprevious", handler.SkipUsersPlaybackToPreviousHandler)
	e.GET("/player/put/seek", handler.SeekToPositionInCurrentlyPlayingTrackHandler)


	// Require SSL
	e.Logger.Fatal(e.StartTLS(":3000", "cert.pem", "key.pem"))
}
