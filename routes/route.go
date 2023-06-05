package routes

import (
	"app-api/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.POST("/register", controller.InsertUserController)
	e.GET("/users", controller.GetUserController)
	e.GET("/user/:id", controller.GetUserByIDController)
	e.PUT("/user/:id", controller.UpdateUserByIDContoller)
	e.DELETE("/user/:id", controller.DeleteUserByIDController)
	return e
}
