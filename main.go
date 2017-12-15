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

	e.GET("/", handler.ExampleHandler)
	e.GET("/callback/", handler.ExampleCallbackHandler)

	//HTTPSを使うのに必要！
	e.Logger.Fatal(e.StartTLS(":3000", "cert.pem", "key.pem"))
}
