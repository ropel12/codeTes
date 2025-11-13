package service

import (
	"api/controller"
	"api/repository"
	"api/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewService() {
	db := repository.NewDB()
	controller := controller.NewController(db)
	e := echo.New()
	e.GET("/getdata", func(c echo.Context) error {
		data, err := controller.GetData()
		if err != nil {
			return utils.ResponseAPI(c, data, http.StatusBadRequest)
		}
		return utils.ResponseAPI(c, data, http.StatusOK)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
