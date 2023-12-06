package routes

import (
	"sosmed/features/users"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo, uc users.Handler) {

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	routeUser(e, uc)
}

func routeUser(e *echo.Echo, uc users.Handler) {
	e.GET("/user/id", uc.GetUserById(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))

	e.POST("/register", uc.Register())
	e.POST("/login", uc.Login())

	e.DELETE("/user/id", uc.DelUserById(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))

	e.PUT("/updateuser", uc.UpdateUser(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
}
