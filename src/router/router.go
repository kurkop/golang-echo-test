package router

import (
	"github.com/kurkop/golang-echo-test/src/api"
	"github.com/kurkop/golang-echo-test/src/api/middlewares"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	// create groups
	adminGroup := e.Group("/admin")
	cookieGroup := e.Group("/cookie")
	jwtGroup := e.Group("/jwt")

	// set all middlewares
	middlewares.SetMainMiddlewares(e)
	middlewares.SetAdminMiddlewares(adminGroup)
	middlewares.SetJwtMiddlewares(jwtGroup)
	middlewares.SetCookieMiddlewares(cookieGroup)

	// set main routes
	api.MainGroup(e)
	// set admin routes
	api.AdminGroup(adminGroup)
	// set cookie routes
	api.CookieGroup(cookieGroup)
	// set jwt routes
	api.JwtGroup(jwtGroup)

	return e
}
