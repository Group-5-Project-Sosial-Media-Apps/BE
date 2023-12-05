package routes

import (
	"sosmed/features/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo, uc users.Handler) {

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	routeUser(e, uc)
}

func routeUser(e *echo.Echo, uc users.Handler) {
	e.GET("/user/id", uc.GetUserById())

	e.POST("/register", uc.Register())
	e.POST("/login", uc.Login())
}
