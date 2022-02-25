package handlers

import (
	"backend/business/services/user"
	"database/sql"

	"github.com/labstack/echo"
)

func Api(db *sql.DB , e *echo.Echo){

	h := userHandler{
		user: user.New(db),
	}
	e.GET("/Users/:id",h.GetData)
	e.GET("/Users/",h.GetData)
	e.POST("/Users/",h.EnterData)
	e.PUT("/Users/:id",h.EditData)
	e.DELETE("/Users/:id",h.DeleteData)
	
}

