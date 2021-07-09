package root

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type rootController struct {
}

func NewRootController() *rootController {
	return &rootController{}
}

func (rc *rootController) Root(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello world!")
}
