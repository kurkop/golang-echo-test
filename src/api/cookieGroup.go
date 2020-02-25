package api

import (
	"github.com/kurkop/golang-echo-test/src/api/handlers"
	"github.com/labstack/echo"
)

func CookieGroup(g *echo.Group) {
	g.GET("/main", handlers.MainCookie)
}
