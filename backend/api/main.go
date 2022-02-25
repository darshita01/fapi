package main

import (
	"backend/additionalFunc"
	"backend/api/handlers"
	"backend/database"
	"fmt"

	"github.com/labstack/echo"
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

	e.Logger.Fatal(e.Start(":8004"))
}
