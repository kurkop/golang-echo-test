package api

import (
	"github.com/kurkop/golang-echo-test/src/api/handlers"
	"github.com/labstack/echo"
)

func MainGroup(e *echo.Echo) {
	e.GET("/login", handlers.Login)

	e.GET("/hi", handlers.Yello)
	e.GET("/cats/:data", handlers.GetCat)
	e.POST("/cats", handlers.AddCat)
	e.POST("/dogs", handlers.AddDog)
	e.POST("/hamsters", handlers.AddHamster)
}
