package middlewares

import (
	"context"
	firebase "firebase.google.com/go"
	"github.com/labstack/echo"
	fa "firebase.google.com/go/auth"
	"google.golang.org/api/option"
	"net/http"
	"strings"
)

type Context struct {
	echo.Context
	Token *fa.Token
}

func Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			rcc := &Context{
				Context: c,
				Token: nil,
			}

			opt := option.WithCredentialsFile("/Users/joshuaobasaju/Documents/firebase.json")
			app, err := firebase.NewApp(context.Background(), nil, opt)

			if err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}

			auth, err := app.Auth(context.Background())
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}

			header := c.Request().Header.Get(echo.HeaderAuthorization)
			idToken := strings.TrimSpace(strings.Replace(header, "Bearer", "", 1))
			token, err := auth.VerifyIDToken(context.Background(), idToken)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}

			rcc.Token = token
			return next(rcc)
		}
	}
}

func WrapCustomContext(fn func(c *Context) error) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return fn(ctx.(*Context))
	}
}