package middlewares

import (
	"net/http"
	"rozhok/utils/helper"

	"github.com/labstack/echo/v4"
)

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, role, _ := ExtractToken(c)
		if role != "admin" {
			return c.JSON(http.StatusForbidden, helper.FailedResponseHelper("role not admin"))
		}
		return next(c)
	}
}

func IsPorter(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, role, _ := ExtractToken(c)
		if role != "porter" {
			return c.JSON(http.StatusForbidden, helper.FailedResponseHelper("role not porter"))
		}
		return next(c)
	}
}

func IsJunkStation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, role, _ := ExtractToken(c)
		if role != "junk_station" {
			return c.JSON(http.StatusForbidden, helper.FailedResponseHelper("role not junk_station"))
		}
		return next(c)
	}
}

func IsJunkStationVerified(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, role, status := ExtractToken(c)
		if role != "junk_station" && status != "terverifikasi" {
			return c.JSON(http.StatusForbidden, helper.FailedResponseHelper("role not junk_station"))
		}
		return next(c)
	}
}

func IsClient(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, role, _ := ExtractToken(c)
		if role != "client" {
			return c.JSON(http.StatusForbidden, helper.FailedResponseHelper("role not client"))
		}
		return next(c)
	}
}
