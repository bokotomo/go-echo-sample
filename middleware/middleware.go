package middleware

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)

func Init (e *echo.Echo) *echo.Echo {
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    return e
}
