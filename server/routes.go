package server

import (
	"firebase-go/pkg/middlewares"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func initRoutes(e *echo.Group) {
	e.GET("/", getHome())
	e.POST("/users", createUserHandler())
}

func initMiddleWares(e *echo.Echo) *echo.Group {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderXRequestedWith, echo.HeaderContentType, echo.HeaderAuthorization},
	}))

	g := e.Group("api", middlewares.Auth())
	return g
}