package routes

import (
	"github.com/ferriantitus/rest-echo/controllers"
	"net/http"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error{
		return c.String(http.StatusOK, "Hello, my name is Ferrian Titus !")
	})

	e.GET("/user", controllers.FetchAllUserOrID)

	e.GET("/user:id", controllers.FetchAllUserOrID)

	e.POST("/user", controllers.StoreUser)

	e.PUT("/user:id", controllers.UpdateUser)

	e.DELETE("/user:id", controllers.DeleteUser)

	return e
}