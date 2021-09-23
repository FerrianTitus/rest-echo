package controllers

import (
	"github.com/ferriantitus/rest-echo/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

//Function untuk mengecek error yang ada pada models fetch all user or id
func FetchAllUserOrID(c echo.Context) error {
	result, err := models.FetchAllUserOrID(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Function untuk mengecek error yang ada pada models store user
func StoreUser(c echo.Context) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	roles := c.FormValue("roles")

	convRoles, err := strconv.Atoi(roles)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.StoreUser(username, email, convRoles)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

//Function untuk mengecek error yang ada pada models update user
func UpdateUser(c echo.Context) error {
	id := c.FormValue("id")
	username := c.FormValue("username")
	email := c.FormValue("email")
	roles := c.FormValue("roles")

	convID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	convRoles, err := strconv.Atoi(roles)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateUser(convID, username, email, convRoles)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

//Function untuk mengecek error yang ada pada models delete user
func DeleteUser(c echo.Context) error {
	id := c.FormValue("id")

	convID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteUser(convID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}