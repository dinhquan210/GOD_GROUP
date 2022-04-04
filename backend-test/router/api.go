package router

import (
	"backend-test/handler"

	"github.com/labstack/echo/v4"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
}

func (api *API) SetupRouter() {
	api.Echo.GET("/users/sign-in", api.UserHandler.HandleSignIn)
	api.Echo.GET("/users/sign-up", api.UserHandler.HandleSignUp)
}
