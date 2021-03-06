package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetAdminMiddlewares(g *echo.Group) {

	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}`,
	}))

	g.Use((middleware.BasicAuth(func(username string, password string, c echo.Context) (bool, error) {
		// check in the DB
		if username == "admin" && password == "admin" {
			return true, nil
		}
		return false, nil
	})))

}
