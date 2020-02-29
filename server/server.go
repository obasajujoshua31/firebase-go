package server

import (
	"fmt"
	"github.com/labstack/echo"
	"log"
	"os"
)


func StartServer() error {
	e := echo.New()

	g := initMiddleWares(e)
	initRoutes(g)

	log.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("APP_HOST"))))
	return nil
}
