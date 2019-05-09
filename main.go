package main

import (
    "github.com/labstack/echo"
	  "./route"
    "./middleware"
)

func main() {
    e := echo.New()
    e = middleware.Init(e)
	  e = route.Init(e)
    e.Start(":8000")
}
