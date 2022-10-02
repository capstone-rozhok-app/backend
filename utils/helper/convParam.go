package helper

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func ParamInt(c echo.Context, paramKey string) int {
	id, err := strconv.Atoi(c.Param(paramKey))
	if err != nil || id < 0 {
		return -1
	}

	return id

}
