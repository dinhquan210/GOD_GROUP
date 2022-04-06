package router

import (
	"backend-test/handler"
	middleware "backend-test/middleware"

	"github.com/labstack/echo/v4"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
	RepoHandler handler.RepoHandler
}

func (api *API) SetupRouter() {

	//user
	api.Echo.POST("/user/sign-in", api.UserHandler.HandleSignIn)
	api.Echo.POST("/user/sign-up", api.UserHandler.HandleSignUp)

	//profile
	user := api.Echo.Group("/user", middleware.JWTMiddleware())
	user.GET("/profile", api.UserHandler.Profile)
	user.PUT("/profile/update", api.UserHandler.UpdateProfile)

	//github
	github := api.Echo.Group("/github", middleware.JWTMiddleware())
	github.GET("/trending", api.RepoHandler.RepoTrending)
}
