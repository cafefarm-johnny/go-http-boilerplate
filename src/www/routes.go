package www

import (
	"github.com/labstack/echo/v4"
	"go-http-boilerplate/src/domain/root"
)

func Register(e *echo.Echo) {
	index(e)
}

func index(e *echo.Echo) {
	rc := root.NewRootController()
	e.GET("/", rc.Root)
}
