package route

import(
	  "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
	  "../controller"
)

func Init(e *echo.Echo) *echo.Echo {
    e.GET("/", controller.Wellcome())
    e.POST("/login", controller.Login())

    auth := e.Group("/users")
    auth.Use(middleware.JWT([]byte("secret")))
    auth.POST("", controller.GetUser())

	  return e
}
