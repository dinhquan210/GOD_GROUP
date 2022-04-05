package router

import (
	"backend-test/handler"
	myMiddleware "backend-test/middleware"

	"github.com/labstack/echo/v4"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
}

func (api *API) SetupRouter() {
	api.Echo.POST("/users/sign-in", api.UserHandler.HandleSignIn, myMiddleware.IsAdmin())
	api.Echo.POST("/users/sign-up", api.UserHandler.HandleSignUp)
}
