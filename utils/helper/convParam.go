package helper

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func ParamInt(c echo.Context) int {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 0 {
		return -1
	}

	return id

}
