package www

import (
	"github.com/labstack/echo/v4"
	"go-http-boilerplate/src/domain/root"
	users2 "go-http-boilerplate/src/domain/users"
)

func Register(e *echo.Echo) {
	index(e)
	users(e)
}

func index(e *echo.Echo) {
	rc := root.NewRootController()
	e.GET("/", rc.Root)
}

func users(e *echo.Echo) {
	uc := users2.NewUserController()
	g := e.Group("/api/v1")
	g.POST("/users/signup", uc.Signup)
	g.POST("/users/signin", uc.Signin)
	g.POST("/users/self-authenticate", uc.SelfAuthenticate)
	g.GET("/users/logout", uc.Logout)
}
