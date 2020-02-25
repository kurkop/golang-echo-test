package handlers

import (
	"net/http"
	"github.com/labstack/echo"
)


func Yello(c echo.Context) error {
	return c.String(http.StatusOK, "Yallo from  from side")
}
