package utils

import (
	"api/domain"
	"fmt"

	"github.com/labstack/echo/v4"
)

func ResponseAPI(c echo.Context, data interface{}, httpCode int) error {
	return c.JSON(httpCode, domain.ResponseApi{
		StatusCode: fmt.Sprintf("%d", httpCode),
		Data:       data,
	})
}
