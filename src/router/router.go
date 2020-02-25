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

	// set routes
	api.MainGroup(e)
	api.AdminGroup(adminGroup)
	api.CookieGroup(cookieGroup)
	api.JwtGroup(jwtGroup)

	return e
}
