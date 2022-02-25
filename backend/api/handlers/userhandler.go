package handlers

import (
	"backend/business/services/user"
	"net/http"

	"github.com/labstack/echo"
)

type userHandler struct{
	user user.Userstruct
}
func (h userHandler) GetData(c echo.Context) error {
	params := c.Param("id")
	response := h.user.GetUserData(params)
	if len(response.Data) == 0 {
		return c.JSON(http.StatusBadRequest, response)
	}
	return c.JSON(http.StatusOK,response)
}

func (h userHandler) EnterData(c echo.Context) error {
	fName := c.FormValue("F_name")
	lName := c.FormValue("L_name")
	mName := c.FormValue("M_name")
	response := h.user.EnterUserData(fName, lName, mName)
	if response.Message == "You have missed some field"{
		return c.JSON(http.StatusBadRequest,response)
	}
	return c.JSON(http.StatusAccepted,response)
}

func (h userHandler) EditData(c echo.Context) error {
	params := c.Param("id")
	fName := c.FormValue("F_name")
	lName := c.FormValue("L_name")
	mName := c.FormValue("M_name")
	response := h.user.EditUserData(fName, lName, mName, params)
	if response.Message == "id is not in table"{
		return c.JSON(http.StatusBadRequest,response)
	}
	return c.JSON(http.StatusOK,response)
}

func (h userHandler) DeleteData(c echo.Context) error {
	params := c.Param("id")
	response := h.user.DeleteUserData(params)
	if response.Message == "id is not in the table"{
		return c.JSON(http.StatusBadRequest,response)
	}
	return c.JSON(http.StatusOK,response)
}