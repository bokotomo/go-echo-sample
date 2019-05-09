package controller

import (
    "net/http"
    "github.com/labstack/echo"
    "github.com/dgrijalva/jwt-go"
    "time"
)

func Login() echo.HandlerFunc {
    return func(c echo.Context) error {
        user := c.FormValue("user")
        password := c.FormValue("password")

        if user == "test" && password == "test" {
            token := jwt.New(jwt.SigningMethodHS256)

            claims := token.Claims.(jwt.MapClaims)
            claims["name"] = "test"
            claims["admin"] = true
            claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

            t, err := token.SignedString([]byte("secret"))
            if err != nil {
                return err
            }
            return c.JSON(http.StatusOK, map[string]string{
                "access_token": t,
            })
        }

        return echo.ErrUnauthorized
    }
}
