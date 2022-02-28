package main

import (
	"backend/additionalFunc"
	"backend/api/handlers"
	"backend/database"
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main(){
	
	db := database.SetupDB()
	defer func() {
		err := db.Close()
		additionalFunc.CheckErr(err)
	}()

	e := echo.New()
	fmt.Println("server at")
	handlers.Api(db , e)

	// headersOk := han.AllowedHeaders([]string{"*"})
	// originsOk := han.AllowedOrigins([]string{"*"})
	// methodsOk := han.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS","DELETE"})
	// e.Use(middleware.CORS(headersOk, originsOk, methodsOk))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	  }))
	e.Logger.Fatal(e.Start(":8004"))
}
