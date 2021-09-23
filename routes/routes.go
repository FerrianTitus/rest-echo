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

	//Routing untuk GET semua data dalam table atau database user
	e.GET("/user", controllers.FetchAllUserOrID)

	//Routing untuk GET data dengan ID
	e.GET("/user:id", controllers.FetchAllUserOrID)

	//Routing untuk POST data ke dalam table atau database user
	e.POST("/user", controllers.StoreUser)

	//Routing untuk melakukan update data
	e.PUT("/user:id", controllers.UpdateUser)

	//Routing untuk DELETE data
	e.DELETE("/user:id", controllers.DeleteUser)

	return e
}