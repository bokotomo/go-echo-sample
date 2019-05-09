package controller

import (
    "net/http"
    "github.com/labstack/echo"
    "github.com/dgrijalva/jwt-go"
)

func GetUser() echo.HandlerFunc  {
    return func(c echo.Context) error {
        user := c.Get("user").(*jwt.Token)
        claims := user.Claims.(jwt.MapClaims)
        name := claims["name"].(string)
        return c.String(http.StatusOK, "Welcom " + name)
    }
}
