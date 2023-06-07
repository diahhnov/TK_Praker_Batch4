package routes

import (
	"app-api/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.POST("/users", controllers.InsertUserController)
	e.GET("/users", controllers.GetUserController)
	e.GET("/user/:id", controllers.GetUserByIDController)
	e.PUT("/user/:id", controllers.UpdateUserByIDContoller)
	e.DELETE("/user/:id", controllers.DeleteUserByIDController)
	e.Use(middleware.Logger())
	return e
}
