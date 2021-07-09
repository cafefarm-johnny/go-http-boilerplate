package main

import (
	"github.com/labstack/echo/v4"
	"go-http-boilerplate/src/www"
)

func main() {
	e := echo.New()
	www.StartUp(e)
}
