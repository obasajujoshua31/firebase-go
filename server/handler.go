package server

import (
	"firebase-go/datastore"
	"firebase-go/domain/model"
	"firebase-go/pkg/middlewares"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/http"
)

func getHome() echo.HandlerFunc {
	return func (c echo.Context) error {
		 return c.String(http.StatusOK, "Welcome")
	}
}

func  createUserHandler() echo.HandlerFunc{
	return middlewares.WrapCustomContext(func(context *middlewares.Context) error {
			var params model.User
			if err := context.Bind(&params); err != nil {
				return context.JSON(http.StatusBadRequest, err.Error())
			}

			db, err := datastore.ConnectToDB()
			if err != nil {
				return context.JSON(http.StatusInternalServerError, err.Error())
			}
			defer db.Close()

			user, err := db.FindUserByUUID(params.UUID)

			if err != nil && !gorm.IsRecordNotFoundError(err)  {
				return context.JSON(http.StatusBadRequest, err.Error())
			}

			if user != nil {
				return context.JSON(http.StatusBadRequest, "User already exist")
			}

			if err := db.CreateUser(&params); err != nil {
				//return context.JSON(http.StatusInternalServerError, err.Error())
				return err
			}

			return context.JSON(http.StatusCreated, &params)
	})
	}